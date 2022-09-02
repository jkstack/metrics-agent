//go:build windows

package user

import "metrics/internal/errors"

func List() ([]User, error) {
	return nil, errors.ErrUnsupported
}
