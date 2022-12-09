package user

import "errors"

// User user object
type User struct {
	Name string
	ID   string
	GID  string
}

var errNotSupported = errors.New("not supported")
