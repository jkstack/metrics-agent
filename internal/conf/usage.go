package conf

import "github.com/jkstack/jkframe/utils"

type usage struct {
	enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
}
