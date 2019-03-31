package main

import "testing"

func Test_run11052Func(t *testing.T){
	if a:=run11052Func(4, "1 5 6 7"); a!=10 {
		t.Error(a)
	}
	if a:=run11052Func(5, "10 9 8 7 6"); a!=50 {
		t.Error(a)
	}
	if a:=run11052Func(10, "1 1 2 3 5 8 13 21 34 55"); a!=55 {
		t.Error(a)
	}
	if a:=run11052Func(10, "5 10 11 12 13 30 35 40 45 47"); a!=50 {
		t.Error(a)
	}
	if a:=run11052Func(4, "5 2 8 10"); a!=20 {
		t.Error(a)
	}
	if a:=run11052Func(4, "3 5 15 16"); a!=18 {
		t.Error(a)
	}
	if a:=run11052Func(10, "1 100 160 1 1 1 1 1 1 1"); a!=520 {
		t.Error(a)
	}

}