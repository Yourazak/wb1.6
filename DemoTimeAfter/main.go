package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var workMS int
	var timeoutMS int
	flag.IntVar(&workMS, "work", 1500, "длительность работы (мс)")
	flag.IntVar(&timeoutMS, "timeout", 1000, "таймаут ожидания (мс)")
	flag.Parse()

	done := make(chan struct{})

	start := time.Now()
	go func() {
		time.Sleep(time.Duration(workMS) * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		fmt.Printf("operation: finished in %v\n", time.Since(start))
	case <-time.After(time.Duration(timeoutMS) * time.Millisecond):
		fmt.Printf("operation: timeout after %v\n", time.Since(start))
	}
	fmt.Println("main: done")
}
