package internal

import (
	"context"
	"metrics/internal/utils"

	"github.com/jkstack/anet"
)

func (agent *Agent) OnConnect() {
	if agent.chWrite != nil {
		close(agent.chWrite)
		agent.chWrite = nil
	}
	agent.chWrite = make(chan *anet.Msg, 10000)
}

func (agent *Agent) OnDisconnect() {
	close(agent.chWrite)
	agent.chWrite = nil
}

func (agent *Agent) OnReportMonitor() {
	defer utils.Recover("report agent status")
	var msg anet.Msg
	msg.Type = anet.TypeHMReportAgentStatus
	msg.HMAgentStatus = &anet.HMAgentStatus{
		Jobs:     agent.cfg.Task.Jobs,
		Warnings: agent.warnings.Load(),
	}
}

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
			}
		}
		agent.OnRewriteConfigure()
		agent.run()
	}
	return nil
}

func (agent *Agent) LoopWrite(ctx context.Context, ch chan *anet.Msg) error {
	for {
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
