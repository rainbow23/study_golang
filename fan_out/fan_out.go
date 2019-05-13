// Package main provides ...
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		tasks, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(tasks) * time.Millisecond
		time.Sleep(d)
		fmt.Println("processing task", tasks)
	}
}
func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}
	close(tasksCh)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}
