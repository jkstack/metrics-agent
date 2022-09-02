//go:build windows
// +build windows

package install

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func Time() (time.Time, error) {
	fi, err := os.Stat(`C:\Windows\system.ini`)
	if err != nil {
		return time.Time{}, err
	}
	st, ok := fi.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return time.Time{}, fmt.Errorf("file sys: %T", fi.Sys())
	}
	ns := st.CreationTime.Nanoseconds()
	return time.Unix(ns/1e9, ns%1e9), nil
}
