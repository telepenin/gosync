package main

import (
	"fmt"
	"sync"
)

func main(){
	r := foo()
	fmt.Println(r)
}

func foo() int {

	var c int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i ++ {
		wg.Add(1)

		go func(){
			mu.Lock()
			c++
			mu.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()

	return c
}
