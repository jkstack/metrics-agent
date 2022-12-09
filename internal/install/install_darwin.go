//go:build darwin
// +build darwin

package install

import (
	"metrics/internal/errors"
	"time"
)

// Time get system install time
func Time() (time.Time, error) {
	return time.Time{}, errors.ErrUnsupported
}
