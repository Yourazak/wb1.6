package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go counter(ctx, &wg)

	<-ctx.Done()
	wg.Wait()
	fmt.Println("main: Done")
}

func counter(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("worker: stop")
			return
		case <-ticker.C:
			i++
			fmt.Println("Counter: ", i)
		}
	}
}
