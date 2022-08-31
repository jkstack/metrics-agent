package conf

import "github.com/jkstack/jkframe/utils"

type usage struct {
	Enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
}
