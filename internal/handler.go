package internal

import (
	"context"

	"github.com/jkstack/anet"
)

func (agent *Agent) OnConnect() {
}

func (agent *Agent) OnDisconnect() {
}

func (agent *Agent) OnReportMonitor() {
}

func (agent *Agent) OnMessage(msg *anet.Msg) error {
	switch msg.Type {
	case anet.TypeHMStaticReq:
		var rep anet.Msg
		rep.Type = anet.TypeHMStaticRep
		rep.TaskID = msg.TaskID
		rep.HMStatic = getStatic()
		agent.chWrite <- &rep
	case anet.TypeHMDynamicReq:
		var rep anet.Msg
		rep.Type = anet.TypeHMDynamicRep
		rep.TaskID = msg.TaskID
		rep.HMDynamicRep = getDynamic(msg.HMDynamicReq, agent.cfg)
		agent.chWrite <- &rep
	case anet.TypeHMQueryCollect:
		var rep anet.Msg
		rep.Type = anet.TypeHMCollectStatus
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
			ch <- msg
		}
	}
}
