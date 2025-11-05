package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		tasks   int
		workers int
		prodMS  int
		procMS  int
		bufSize int
	)
	flag.IntVar(&tasks, "tasks", 20, "сколько задач отправить")
	flag.IntVar(&workers, "workers", 2, "сколько потребителей запустить")
	flag.IntVar(&prodMS, "prod", 120, "интервал генерации задач (мс)")
	flag.IntVar(&procMS, "proc", 180, "время обработки задачи (мс)")
	flag.IntVar(&bufSize, "buf", 4, "размер буфера рабочего канала")
	flag.Parse()

	work := make(chan int, bufSize)
	var wg sync.WaitGroup

	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for jobs := range work {
				fmt.Printf("worker %d: got task %d\n", id, jobs)
				time.Sleep(time.Duration(procMS) * time.Millisecond)
			}
			fmt.Printf("worker %d: channel closed exit\n", id)
		}(w)
	}

	go func() {
		for i := 1; i <= tasks; i++ {
			work <- i
			fmt.Printf("producer: sent task %d\n", i)
			time.Sleep(time.Duration(prodMS) * time.Millisecond)
		}
		close(work)
		fmt.Println("producer: close worlk channel")
	}()

	wg.Wait()
	fmt.Println("main: done")
}
