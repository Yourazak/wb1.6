package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var seconds int
	var tickMS int
	flag.IntVar(&seconds, "seconds", 2, "сколько работать до остановки (сек)")
	flag.IntVar(&tickMS, "tick", 200, "шаг тика (мс)")
	flag.Parse()

	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(time.Duration(tickMS) * time.Millisecond)
		defer ticker.Stop()

		step := 0
		for {
			select {
			case <-done:
				fmt.Println("worker: received done, exit")
				return
			case <-ticker.C:
				step++
				fmt.Println("Step:", step)
			}
		}
	}()

	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("main: closing done")
	close(done)

	wg.Wait()
	fmt.Println("main: done")
}
