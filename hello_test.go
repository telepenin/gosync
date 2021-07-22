package main

import (
	"testing"
	"fmt"
)

func TestMain(t *testing.T){
	var actual int = 1000
	r := foo()
	if r != actual {
		t.Error(fmt.Sprintf("expected, actual: %d vs %d", r, actual))
	}
}
func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		foo()
	}
}