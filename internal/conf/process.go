package conf

import "github.com/jkstack/jkframe/utils"

// ProcessConfigure process configure
type ProcessConfigure struct {
	Enabled  bool           `kv:"-"`
	Limit    int            `kv:"qps_limit"`
	Interval utils.Duration `kv:"interval"`
}
