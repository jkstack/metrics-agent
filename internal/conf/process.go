package conf

import "github.com/jkstack/jkframe/utils"

type process struct {
	enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
	Top      uint64         `kv:"top"`
}
