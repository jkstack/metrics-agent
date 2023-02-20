//go:build !windows
// +build !windows

package dns

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func NameServer() ([]string, error) {
	f, err := os.Open("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var ret []string
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return ret, nil
			}
			return nil, err
		}
		cols := strings.Fields(line)
		if len(cols) < 2 {
			continue
		}
		if cols[0] == "nameserver" {
			ret = append(ret, cols[1])
		}
	}
}
