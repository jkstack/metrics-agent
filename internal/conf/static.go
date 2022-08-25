package conf

import "github.com/jkstack/jkframe/utils"

type static struct {
	Interval utils.Duration `kv:"interval"`
	enabled  bool           `kv:"-"`
}
