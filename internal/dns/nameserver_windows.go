//go:build windows

package dns

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func NameServer() ([]string, error) {
	var b []byte
	l := uint32(15000)
	for {
		b = make([]byte, l)
		err := windows.GetAdaptersAddresses(syscall.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l)
		if err == nil {
			if l == 0 {
				return nil, err
			}
			break
		}
		if err.(syscall.Errno) != syscall.ERROR_BUFFER_OVERFLOW {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
		if l <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
	}
	addrs := make(map[string]bool)
	for adapter := (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])); adapter != nil; adapter = adapter.Next {
		address := adapter.FirstDnsServerAddress
		for address != nil {
			addrs[address.Address.IP().String()] = true
			address = address.Next
		}
	}
	ret := make([]string, 0, len(addrs))
	for addr := range addrs {
		ret = append(ret, addr)
	}
	return ret, nil
}
