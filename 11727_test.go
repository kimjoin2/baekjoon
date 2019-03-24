package main

import "testing"

func Test_run11727Func(t *testing.T){
	if 3 != run11727Func(2) {
		t.Error()
	}
	if 171 != run11727Func(8) {
		t.Error()
	}
	if 2731 != run11727Func(12) {
		t.Error()
	}
}