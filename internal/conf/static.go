package conf

import "github.com/jkstack/jkframe/utils"

type static struct {
	Enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
}
