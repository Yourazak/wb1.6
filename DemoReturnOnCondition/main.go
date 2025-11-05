package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var limit int
	var tickMS int
	flag.IntVar(&limit, "limit", 7, "после какого шага выйти return")
	flag.IntVar(&tickMS, "tick", 200, "шаг тика (мс)")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(time.Duration(tickMS) * time.Millisecond)
		defer ticker.Stop()

		for step := 1; ; step++ {
			<-ticker.C
			fmt.Println("Step:", step)
			if step >= limit {
				fmt.Println("worker: condition met, return")
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println("main: done")
}
