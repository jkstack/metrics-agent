package conf

import "github.com/jkstack/jkframe/utils"

type TempsConfigure struct {
	Enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
}
