//go:build windows

package user

func List() ([]User, error) {
	return nil, errNotSupported
}
