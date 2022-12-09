//go:build windows

package user

import "metrics/internal/errors"

// List get system user list
func List() ([]User, error) {
	return nil, errors.ErrUnsupported
}
