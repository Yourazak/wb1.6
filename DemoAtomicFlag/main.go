package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var seconds int
	flag.IntVar(&seconds, "seconds", 1, "how long to run before stopping (seconds)")
	var wg sync.WaitGroup
	var stop atomic.Bool

	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		step := 0
		for {
			if stop.Load() {
				fmt.Println("worker: stop flag observed, exiting")
				return
			}
			select {
			case <-ticker.C:
				step++
				fmt.Println("Step: ", step)
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Get stop flag")
	stop.Store(true)

	wg.Wait()
	fmt.Println("Done")
}
