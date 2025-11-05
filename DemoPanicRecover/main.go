package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var failAt int
	flag.IntVar(&failAt, "fail", 5, "на каком шаге сымитировать панику")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg, failAt)

	wg.Wait()
	fmt.Println("main: done")
}

func worker(wg *sync.WaitGroup, failAt int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("worker: recovered from panic:", r)
		}
		wg.Done()
	}()

	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	step := 0
	for {
		select {
		case <-ticker.C:
			step++
			fmt.Println("Step:", step)
			if step == failAt {
				panic(fmt.Sprintf("boom at step %d", step))
			}
		}
	}
}
