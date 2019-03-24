package main

import "testing"

func Test_run11053Func(t *testing.T){
	if a := run11053Func(6, "10 20 10 30 20 50"); a != 4 {
		t.Error(a)
	}
	if a := run11053Func(7, "8 6 9 1 4 6 7"); a != 4 {
		t.Error(a)
	}
	if a := run11053Func(10, "4 3 7 4 7 2 5 2 10 7"); a != 4 {
		t.Error(a)
	}
	if a := run11053Func(8, "2 6 9 5 6 3 8 8"); a != 4 {
		t.Error(a)
	}
	if a := run11053Func(6, "6 9 4 5 3 7"); a != 3 {
		t.Error(a)
	}
	if a := run11053Func(6, "5 2 4 8 7 2"); a != 3 {
		t.Error(a)
	}
}
