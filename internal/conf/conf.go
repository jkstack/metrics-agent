package conf

import (
	"strings"

	"github.com/jkstack/libagent/conf"
)

type task struct {
	Jobs    stringList       `kv:"jobs"`
	Static  static           `kv:"static"`
	Usage   usage            `kv:"usage"`
	Process ProcessConfigure `kv:"process"`
	Conns   ConnsConfigure   `kv:"conns"`
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
			task.Static.Enabled = true
		case "usage":
			task.Usage.Enabled = true
		case "process":
			task.Process.Enabled = true
		case "conns":
			task.Conns.Enabled = true
		}
	}
}
