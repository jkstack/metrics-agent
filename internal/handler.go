package internal

import (
	"context"
	"metrics/internal/utils"
	"sync/atomic"
	"time"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
)

// OnConnect on connect callback
func (agent *Agent) OnConnect() {
	if agent.chWrite != nil {
		close(agent.chWrite)
		agent.chWrite = nil
	}
	agent.chWrite = make(chan *anet.Msg, 10000)
}

// OnDisconnect on disconnect callback
func (agent *Agent) OnDisconnect() {
	if agent.chWrite != nil {
		close(agent.chWrite)
	}
	agent.chWrite = nil
}

// OnReportMonitor on report monitor callback
func (agent *Agent) OnReportMonitor() {
	defer utils.Recover("report agent status")
	var msg anet.Msg
	msg.Type = anet.TypeHMReportAgentStatus
	var jobs []anet.HMAgentJob
	for _, job := range agent.cfg.Task.Jobs {
		var interval, sent, count uint64
		switch job {
		case "static":
			interval = uint64(agent.cfg.Task.Static.Interval.Duration().Seconds())
			sent = agent.tkStatic.bytes
			count = agent.tkStatic.count
		case "usage":
			interval = uint64(agent.cfg.Task.Usage.Interval.Duration().Seconds())
			sent = agent.tkUsage.bytes
			count = agent.tkUsage.count
		case "process":
			interval = uint64(agent.cfg.Task.Process.Interval.Duration().Seconds())
			sent = agent.tkProcess.bytes
			count = agent.tkProcess.count
		case "conns":
			interval = uint64(agent.cfg.Task.Conns.Interval.Duration().Seconds())
			sent = agent.tkConns.bytes
			count = agent.tkConns.count
		case "temps":
			interval = uint64(agent.cfg.Task.Temps.Interval.Duration().Seconds())
			sent = agent.tkTemps.bytes
			count = agent.tkTemps.count
		}
		jobs = append(jobs, anet.HMAgentJob{
			Name:      job,
			Interval:  interval,
			BytesSent: sent,
			Count:     count,
		})
	}
	msg.HMAgentStatus = &anet.HMAgentStatus{
		Jobs:     jobs,
		Warnings: atomic.LoadUint64(&agent.warnings),
	}
	agent.chWrite <- &msg
}

// OnMessage on receive message callback
func (agent *Agent) OnMessage(msg *anet.Msg) error {
	defer utils.Recover("OnMessage")
	switch msg.Type {
	case anet.TypeHMStaticReq:
		var rep anet.Msg
		rep.Type = anet.TypeHMStaticRep
		rep.TaskID = msg.TaskID
		rep.HMStatic = getStatic(&agent.warnings)
		agent.chWrite <- &rep
	case anet.TypeHMDynamicReq:
		var rep anet.Msg
		rep.Type = anet.TypeHMDynamicRep
		rep.TaskID = msg.TaskID
		rep.HMDynamicRep = getDynamic(msg.HMDynamicReq, agent.cfg, &agent.warnings)
		agent.chWrite <- &rep
	case anet.TypeHMQueryCollect:
		var rep anet.Msg
		rep.Type = anet.TypeHMCollectStatus
		rep.TaskID = msg.TaskID
		var jobs []anet.HMJob
		if agent.cfg.Task.Static.Enabled {
			jobs = append(jobs, anet.HMJob{
				Name:     "static",
				Interval: agent.cfg.Task.Static.Interval.Duration(),
			})
		}
		if agent.cfg.Task.Usage.Enabled {
			jobs = append(jobs, anet.HMJob{
				Name:     "usage",
				Interval: agent.cfg.Task.Usage.Interval.Duration(),
			})
		}
		if agent.cfg.Task.Process.Enabled {
			jobs = append(jobs, anet.HMJob{
				Name:     "process",
				Interval: agent.cfg.Task.Process.Interval.Duration(),
			})
		}
		if agent.cfg.Task.Conns.Enabled {
			jobs = append(jobs, anet.HMJob{
				Name:     "conns",
				Interval: agent.cfg.Task.Conns.Interval.Duration(),
			})
		}
		if agent.cfg.Task.Temps.Enabled {
			jobs = append(jobs, anet.HMJob{
				Name:     "temps",
				Interval: agent.cfg.Task.Temps.Interval.Duration(),
			})
		}
		rep.HMCollectStatus = &anet.HMReportStatusPayload{
			Jobs:       jobs,
			ConnsAllow: agent.cfg.Task.Conns.Allow,
		}
		agent.chWrite <- &rep
	case anet.TypeHMChangeCollectStatus:
		agent.cfg.Task.Static.Enabled = false
		agent.cfg.Task.Usage.Enabled = false
		agent.cfg.Task.Process.Enabled = false
		agent.cfg.Task.Conns.Enabled = false
		agent.cfg.Task.Temps.Enabled = false
		for _, job := range msg.HMChangeStatus.Jobs {
			switch job {
			case "static":
				agent.cfg.Task.Static.Enabled = true
			case "usage":
				agent.cfg.Task.Usage.Enabled = true
			case "process":
				agent.cfg.Task.Process.Enabled = true
			case "conns":
				agent.cfg.Task.Conns.Enabled = true
			case "temps":
				agent.cfg.Task.Temps.Enabled = true
			}
		}
		agent.OnRewriteConfigure()
		agent.run()
		logging.Info("jobs: %v", msg.HMChangeStatus.Jobs)
	}
	return nil
}

// LoopWrite loop write
func (agent *Agent) LoopWrite(ctx context.Context, ch chan *anet.Msg) error {
	for {
		if agent.chWrite == nil {
			time.Sleep(time.Second)
			continue
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-agent.chWrite:
			if msg == nil {
				continue
			}
			ch <- msg
		}
	}
}
