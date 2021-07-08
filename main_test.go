package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	out := numIslands(in)
	want := 1
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}
