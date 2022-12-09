package conf

import "github.com/jkstack/jkframe/utils"

// TempsConfigure temps configure
type TempsConfigure struct {
	Enabled  bool           `kv:"-"`
	Interval utils.Duration `kv:"interval"`
}
