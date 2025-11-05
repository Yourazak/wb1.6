package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 1500*time.Millisecond, "через сколько прервать контекст")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()

		step := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker: context canceled:", ctx.Err())
				return
			case <-ticker.C:
				step++
				fmt.Println("Step:", step)
			}
		}
	}()

	<-ctx.Done()
	wg.Wait()
	fmt.Println("main: Done")
}
