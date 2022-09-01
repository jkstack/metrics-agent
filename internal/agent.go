package internal

import (
	"context"
	"metrics/internal/conf"
	"metrics/internal/utils"
	"sync/atomic"
	"time"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
)

type Agent struct {
	cfgDir  string
	cfg     *conf.Configure
	version string
	chWrite chan *anet.Msg
	// runtime
	ctx    context.Context
	cancel context.CancelFunc
	// monitor
	warnings atomic.Uint64
}

func New(dir, version string) *Agent {
	ag := &Agent{
		cfgDir:  dir,
		cfg:     load(dir),
		version: version,
		chWrite: make(chan *anet.Msg),
	}
	go ag.run()
	return ag
}

func (agent *Agent) AgentName() string {
	return "metrics-agent"
}

func (agent *Agent) Version() string {
	return agent.version
}

func (agent *Agent) run() {
	if agent.cancel != nil {
		agent.cancel()
	}
	agent.ctx, agent.cancel = context.WithCancel(context.Background())
	run := func(interval time.Duration, cb func() *anet.Msg) {
		defer utils.Recover("report")
		for {
			select {
			case <-agent.ctx.Done():
				return
			case <-time.After(interval):
				agent.chWrite <- cb()
			}
		}
	}
	if agent.cfg.Task.Static.Enabled {
		go run(agent.cfg.Task.Static.Interval.Duration(), func() *anet.Msg {
			logging.Info("report static info")
			var msg anet.Msg
			msg.Type = anet.TypeHMStaticRep
			msg.HMStatic = getStatic(&agent.warnings)
			return &msg
		})
	}
	if agent.cfg.Task.Usage.Enabled {
		go run(agent.cfg.Task.Usage.Interval.Duration(), func() *anet.Msg {
			logging.Info("report usage info")
			var msg anet.Msg
			msg.Type = anet.TypeHMDynamicRep
			msg.HMDynamicRep = &anet.HMDynamicRep{
				Usage: getUsage(&agent.warnings),
			}
			return &msg
		})
	}
	if agent.cfg.Task.Process.Enabled {
		go run(agent.cfg.Task.Process.Interval.Duration(), func() *anet.Msg {
			logging.Info("report process list")
			var msg anet.Msg
			msg.Type = anet.TypeHMDynamicRep
			msg.HMDynamicRep = &anet.HMDynamicRep{
				Process: getProcessList(agent.cfg.Task.Process, &agent.warnings, 0),
			}
			return &msg
		})
	}
	if agent.cfg.Task.Conns.Enabled {
		go run(agent.cfg.Task.Conns.Interval.Duration(), func() *anet.Msg {
			logging.Info("report connections list")
			var msg anet.Msg
			msg.Type = anet.TypeHMDynamicRep
			msg.HMDynamicRep = &anet.HMDynamicRep{
				Connections: getConnectionList(agent.cfg.Task.Conns, &agent.warnings, agent.cfg.Task.Conns.Allow),
			}
			return &msg
		})
	}
}
