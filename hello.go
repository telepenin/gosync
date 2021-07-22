package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

func main(){
	r := foo()
	fmt.Println(r)
}

func foo() int {

	var c int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i ++ {
		wg.Add(1)

		func(){
			mu.Lock()
			c++
			mu.Unlock()
			// atomic.AddInt32(&c, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	return c
	// return atomic.LoadInt32(&c)
}
