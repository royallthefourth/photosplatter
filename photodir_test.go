package main

import "testing"

func Test_readPath(t *testing.T) {
	_, err := readPath("")
	if err == nil {
		t.Errorf("Expected error on empty path but got success")
	}
}
