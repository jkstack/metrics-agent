package conf

import "github.com/jkstack/jkframe/utils"

type shell struct {
	enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
	Scripts  stringList     `kv:"scripts"`
	Timeout  utils.Duration `kv:"timeout"`
}
