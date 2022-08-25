package conf

import "github.com/jkstack/jkframe/utils"

type conns struct {
	enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
	Allow    stringList     `kv:"allow"`
}
