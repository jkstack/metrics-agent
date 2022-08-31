//go:build darwin
// +build darwin

package internal

import (
	"time"
)

func getInstallTime() (time.Time, error) {
	return time.Time{}, errUnsupported
}
