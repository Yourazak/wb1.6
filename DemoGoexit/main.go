package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer func() {
			fmt.Println("worker: deferred cleanup before Goexit")
			wg.Done()
		}()

		for i := 1; i <= 10; i++ {
			fmt.Println("worker step:", i)
			time.Sleep(120 * time.Millisecond)
			if i == 5 {
				fmt.Println("worker: calling runtime.Goexit()")
				runtime.Goexit()
			}
		}
	}()

	wg.Wait()
	fmt.Println("main: done")
}
