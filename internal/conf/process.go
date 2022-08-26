package conf

import "github.com/jkstack/jkframe/utils"

type ProcessConfigure struct {
	enabled  bool           `kv:"-"`
	Limit    int            `kv:"qps_limit"`
	Interval utils.Duration `kv:"interval"`
}
