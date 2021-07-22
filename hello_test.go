package main

import (
	"testing"
	"fmt"
)

func TestMain(t *testing.T){
	var actual int32 = 1000
	r := foo()
	if r != actual {
		t.Error(fmt.Sprintf("expected, actual: %d vs %d", r, actual))
	}
}