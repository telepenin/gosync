package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	r := foo()
	fmt.Println(r)
}

func foo() int32 {

	var c int32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i ++ {
		wg.Add(1)

		go func(){
			atomic.AddInt32(&c, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	return atomic.LoadInt32(&c)
}
