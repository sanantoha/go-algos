package main

import (
	"reflect"
	"testing"
)

func TestValidIpAddress(t *testing.T) {
	str := "1921680"

	exp := []string{
		"1.9.216.80",
		"1.92.16.80",
		"1.92.168.0",
		"19.2.16.80",
		"19.2.168.0",
		"19.21.6.80",
		"19.21.68.0",
		"19.216.8.0",
		"192.1.6.80",
		"192.1.68.0",
		"192.16.8.0",
	}

	res := validIPAddresses(str)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("TestValidIpAddress failed, exp: %v, got: %v", exp, res)
	}
}
