package conf

import "github.com/jkstack/jkframe/utils"

type net struct {
	enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
	Allow    stringList     `kv:"allow"`
	Filter   stringList     `kv:"filter"`
}
