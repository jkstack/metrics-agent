package conf

import "github.com/jkstack/jkframe/utils"

type ProcessConfigure struct {
	Enabled  bool           `kv:"-"`
	Limit    int            `kv:"qps_limit"`
	Interval utils.Duration `kv:"interval"`
}
