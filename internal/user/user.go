package user

import "errors"

type User struct {
	Name string
	ID   string
	GID  string
}

var errNotSupported = errors.New("not supported")
