//go:build linux
// +build linux

package install

import (
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func Time() (time.Time, error) {
	files, err := filepath.Glob("/etc/*")
	if err != nil {
		return time.Time{}, err
	}
	t := time.Now()
	for _, file := range files {
		fi, err := os.Stat(file)
		if err != nil {
			continue
		}
		st, ok := fi.Sys().(*syscall.Stat_t)
		if !ok {
			continue
		}
		ts := time.Unix(int64(st.Ctim.Sec), int64(st.Ctim.Nsec))
		if ts.Before(t) {
			t = ts
		}
	}
	return t, nil
}
