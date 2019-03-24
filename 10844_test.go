package main

import "testing"

func Test_run10844Func(t *testing.T){
	if 9 != run10844Func(1) {
		t.Error()
	}

	if 17 != run10844Func(2) {
		t.Error()
	}
	t.Log(run10844Func(100))
}
