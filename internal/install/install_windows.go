//go:build windows
// +build windows

package install

import (
	"time"

	"github.com/yusufpapurcu/wmi"
)

func Time() (time.Time, error) {
	type Win32_OperatingSystem struct {
		InstallDate time.Time
	}
	var list []Win32_OperatingSystem
	q := wmi.CreateQuery(&list, "")
	err := wmi.Query(q, &list)
	if err != nil {
		return time.Time{}, err
	}
	return list[0].InstallDate, nil
}
