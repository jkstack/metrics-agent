//go:build unix || linux
// +build unix linux

package user

import (
	"bufio"
	"os"
	"os/user"
	"strings"
)

func List() ([]User, error) {
	f, err := os.Open("/etc/passwd")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ret []User
	s := bufio.NewScanner(f)
	for s.Scan() {
		if equal := strings.Index(s.Text(), "#"); equal >= 0 {
			continue
		}
		lineSlice := strings.FieldsFunc(s.Text(), func(divide rune) bool {
			return divide == ':'
		})

		if len(lineSlice) > 0 {
			u, err := user.Lookup(lineSlice[0])
			if err != nil {
				continue
			}
			ret = append(ret, User{
				Name: u.Name,
				ID:   u.Uid,
				GID:  u.Gid,
			})
		}

		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
