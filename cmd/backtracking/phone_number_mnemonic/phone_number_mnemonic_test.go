package main

import (
	"reflect"
	"testing"
)

func TestPhoneNumberMnemonic(t *testing.T) {

	exp := []string{"1w0j", "1w0k", "1w0l", "1x0j", "1x0k", "1x0l", "1y0j", "1y0k", "1y0l", "1z0j", "1z0k", "1z0l"}

	res := phoneNumberMnemonics("1905")
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
