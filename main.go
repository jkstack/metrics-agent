package main

import (
	"flag"
	"fmt"
	"metrics/internal"
	"os"
	"path/filepath"
	rt "runtime"

	"github.com/jkstack/jkframe/utils"
	agent "github.com/jkstack/libagent"
)

var (
	version      string = "0.0.0"
	gitBranch    string = "<branch>"
	gitHash      string = "<hash>"
	gitReversion string = "0"
	buildTime    string = "0000-00-00 00:00:00"
)

func showVersion() {
	fmt.Printf("version: %s\ncode version: %s.%s.%s\nbuild time: %s\ngo version: %s\n",
		version,
		gitBranch, gitHash, gitReversion,
		buildTime,
		rt.Version())
}

func main() {
	cf := flag.String("conf", "", "config file dir")
	ver := flag.Bool("version", false, "show version info")
	act := flag.String("action", "", "install, uninstall, start, stop")
	flag.Parse()

	if *ver {
		showVersion()
		return
	}

	if len(*cf) == 0 {
		fmt.Println("missing -conf argument")
		os.Exit(1)
	}

	var app agent.App

	dir, err := filepath.Abs(*cf)
	utils.Assert(err)

	dummy := agent.NewDummyApp("metrics-agent", dir)

	switch *act {
	case "install":
		agent.RegisterService(dummy)
	case "uninstall":
		agent.UnregisterService(dummy)
	case "start":
		agent.Start(dummy)
	case "stop":
		agent.Stop(dummy)
	default:
		app = internal.New(dir, version)
		agent.Run(app)
	}
}
