package core

import (
	"metrics/core/conf"

	"github.com/jkstack/anet"
)

type Agent struct {
	cfgDir  string
	cfg     *conf.Configure
	version string
	chWrite chan *anet.Msg
}

func New(dir, version string) *Agent {
	return &Agent{
		cfgDir:  dir,
		cfg:     load(dir),
		version: version,
		chWrite: make(chan *anet.Msg, 10000),
	}
}

func (agent *Agent) AgentName() string {
	return "metrics-agent"
}

func (agent *Agent) Version() string {
	return agent.version
}
