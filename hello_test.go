package main

import (
	"testing"
	"fmt"
)

func test_main(t testing.T){
	actual := 1000
	r := foo()
	if r != actual {
		t.Error(fmt.Sprintf("expected, actual: %d vs %d", r, actual))
	}
}