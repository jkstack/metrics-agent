package conf

import "github.com/jkstack/jkframe/utils"

// ConnsConfigure connection configure
type ConnsConfigure struct {
	Enabled  bool           `kv:"-"`
	Limit    int            `kv:"qps_limit"`
	Interval utils.Duration `kv:"interval"`
	Allow    stringList     `kv:"allow"`
}
