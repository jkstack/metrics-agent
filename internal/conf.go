package internal

import (
	lconf "metrics/internal/conf"
	"os"

	"github.com/jkstack/jkframe/conf/kvconf"
	"github.com/jkstack/jkframe/logging"
	"github.com/jkstack/jkframe/utils"
	"github.com/jkstack/libagent/conf"
)

func load(dir string) *lconf.Configure {
	f, err := os.Open(dir)
	utils.Assert(err)
	defer f.Close()
	var ret lconf.Configure
	utils.Assert(kvconf.NewDecoder(f).Decode(&ret))
	if ret.Task.Process.Limit < 0 {
		ret.Task.Process.Limit = 200
	}
	if ret.Task.Conns.Limit < 0 {
		ret.Task.Conns.Limit = 500
	}
	for _, task := range ret.Task.Jobs {
		switch task {
		case "static":
			ret.Task.Static.Enabled = true
		case "usage":
			ret.Task.Usage.Enabled = true
		case "process":
			ret.Task.Process.Enabled = true
		case "conns":
			ret.Task.Conns.Enabled = true
		}
	}
	logging.Info("jobs: %v", ret.Task.Jobs)
	return &ret
}

func (agent *Agent) ConfDir() string {
	return agent.cfgDir
}

func (agent *Agent) Configure() *conf.Configure {
	return &agent.cfg.Basic
}

func (agent *Agent) OnRewriteConfigure() error {
	f, err := os.Create(agent.cfgDir + ".tmp")
	if err != nil {
		return err
	}
	defer f.Close()
	defer os.Remove(f.Name())
	err = kvconf.NewEncoder(f).Encode(agent.cfg)
	if err != nil {
		return err
	}
	f.Close()
	os.Rename(agent.cfgDir+".tmp", agent.cfgDir)
	return nil
}
