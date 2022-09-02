package internal

import (
	"context"
	"encoding/json"
	"metrics/internal/conf"
	"metrics/internal/utils"
	"sync/atomic"
	"time"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
)

type tick struct {
	bytes uint64
	count uint64
}

type Agent struct {
	cfgDir  string
	cfg     *conf.Configure
	version string
	chWrite chan *anet.Msg
	// runtime
	ctx    context.Context
	cancel context.CancelFunc
	// monitor
	warnings  uint64
	tkStatic  tick
	tkUsage   tick
	tkProcess tick
	tkConns   tick
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
	run := func(interval time.Duration, tk *tick, cb func() *anet.Msg) {
		save := func() {
			defer utils.Recover("save")
			msg := cb()
			if agent.chWrite == nil {
				return
			}
			agent.chWrite <- msg
			data, _ := json.Marshal(msg)
			atomic.AddUint64(&tk.bytes, uint64(len(data)))
			atomic.AddUint64(&tk.count, 1)
		}
		save()
		for {
			select {
			case <-agent.ctx.Done():
				return
			case <-time.After(interval):
				save()
			}
		}
	}
	if agent.cfg.Task.Static.Enabled {
		go run(agent.cfg.Task.Static.Interval.Duration(), &agent.tkStatic, func() *anet.Msg {
			logging.Info("report static info")
			var msg anet.Msg
			msg.Type = anet.TypeHMStaticRep
			msg.HMStatic = getStatic(&agent.warnings)
			return &msg
		})
	}
	if agent.cfg.Task.Usage.Enabled {
		go run(agent.cfg.Task.Usage.Interval.Duration(), &agent.tkUsage, func() *anet.Msg {
			logging.Info("report usage info")
			var msg anet.Msg
			msg.Type = anet.TypeHMDynamicRep
			begin := time.Now()
			usage := getUsage(&agent.warnings)
			end := time.Now()
			msg.HMDynamicRep = &anet.HMDynamicRep{
				Begin: begin,
				End:   end,
				Usage: usage,
			}
			return &msg
		})
	}
	if agent.cfg.Task.Process.Enabled {
		go run(agent.cfg.Task.Process.Interval.Duration(), &agent.tkProcess, func() *anet.Msg {
			logging.Info("report process list")
			var msg anet.Msg
			msg.Type = anet.TypeHMDynamicRep
			begin := time.Now()
			process := getProcessList(agent.cfg.Task.Process, &agent.warnings, 0)
			end := time.Now()
			msg.HMDynamicRep = &anet.HMDynamicRep{
				Begin:   begin,
				End:     end,
				Process: process,
			}
			return &msg
		})
	}
	if agent.cfg.Task.Conns.Enabled {
		go run(agent.cfg.Task.Conns.Interval.Duration(), &agent.tkConns, func() *anet.Msg {
			logging.Info("report connections list")
			var msg anet.Msg
			msg.Type = anet.TypeHMDynamicRep
			begin := time.Now()
			conns := getConnectionList(agent.cfg.Task.Conns, &agent.warnings, agent.cfg.Task.Conns.Allow)
			end := time.Now()
			msg.HMDynamicRep = &anet.HMDynamicRep{
				Begin:       begin,
				End:         end,
				Connections: conns,
			}
			return &msg
		})
	}
}
