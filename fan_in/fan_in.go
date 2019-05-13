// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
	}
}
func reader(out chan int) {
	for i := range out {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	out := make(chan int)
	go producer(ch, 100*time.Millisecond)
	go producer(ch, 250*time.Millisecond)
	go reader(out)

	for i := range ch {
		out <- i
	}
}
