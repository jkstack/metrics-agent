package conf

import (
	"strings"

	"github.com/jkstack/libagent/conf"
)

type task struct {
	Jobs    stringList `kv:"jobs"`
	Static  static     `kv:"static"`
	Usage   usage      `kv:"usage"`
	Net     net        `kv:"net"`
	Process process    `kv:"process"`
	Conns   conns      `kv:"conns"`
	Shell   shell      `kv:"shell"`
}

// Configure configure
type Configure struct {
	Basic conf.Configure `kv:"basic"`
	Task  task           `kv:"task"`
}

type stringList []string

// MarshalKV marshal jobs
func (sl stringList) MarshalKV() (string, error) {
	return strings.Join(sl, ","), nil
}

// UnmarshalKV unmarshal duration
func (sl *stringList) UnmarshalKV(value string) error {
	tmp := strings.Split(value, ",")
	if len(tmp) == 1 && tmp[0] == "" {
		tmp = []string{}
	}
	*sl = tmp
	return nil
}

// Parse parse jobs and set enabled
func (task *task) Parse() {
	for _, job := range task.Jobs {
		switch job {
		case "static":
			task.Static.enabled = true
		case "usage":
			task.Usage.enabled = true
		case "process":
			task.Process.enabled = true
		case "net":
			task.Net.enabled = true
		case "conns":
			task.Conns.enabled = true
		case "shell":
			task.Shell.enabled = true
		}
	}
}
