// Package main provides ...
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)

	/*
	 * boom := time.After(100 * time.Millisecond)
	 * tick := time.Tick(600 * time.Millisecond)
	 */
Outerloop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break Outerloop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("############################")
}
