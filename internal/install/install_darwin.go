//go:build darwin
// +build darwin

package install

import (
	"metrics/internal/errors"
	"time"
)

func Time() (time.Time, error) {
	return time.Time{}, errors.ErrUnsupported
}
