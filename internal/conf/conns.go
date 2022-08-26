package conf

import "github.com/jkstack/jkframe/utils"

type conns struct {
	enabled  bool           `kv:"-"`
	Limit    int            `kv:"qps_limit"`
	Interval utils.Duration `kv:"interval"`
	Allow    stringList     `kv:"allow"`
}
