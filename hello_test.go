package main

import (
	"testing"
	"fmt"
)

func TestMain(t *testing.T){
	actual := 1000
	r := foo()
	if r != actual {
		t.Error(fmt.Sprintf("expected, actual: %d vs %d", r, actual))
	}
}