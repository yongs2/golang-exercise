package main

import (
	"net"
	"testing"
)

func checkIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}
	return true
}

func TestValidIpAddress(t *testing.T) {
	testCases1 := []struct {
		name string
		ip   string
		want bool
	}{
		{"ValidIPv4", "10.40.210.253", true},
		{"InvalidIPv4", "1000.40.210.253", false},
		{"ValidIPv6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"InvalidIPv6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334:3445", false},
	}

	for _, tc := range testCases1 {
		t.Run(tc.name, func(t *testing.T) {
			if err := checkIPAddress(tc.ip); err != tc.want {
				t.Errorf("got:%v, want:%v\n", err, tc.want)
			}
		})
	}
}
