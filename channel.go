package main

import (
	"fmt"
	"time"
)

// func listenChannel(myCh chan int) {
//     for out := range myCh {
//         fmt.Println(
//             fmt.Sprintf("out form channel => %d", out),
//         )
//     }
// }

// func main() {
//     fmt.Println("This is func main")

//     myChan := make(chan int)

//     go listenChannel(myChan)

//     myChan <- 100
//     myChan <- 200

//     time.Sleep(1 * time.Second) //mainを延命する

//     fmt.Printf("main exited now\n")
// }

func repeat(say string, count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(say, fmt.Sprintf("%d回目", i))
	}
}

func main() {
	go repeat("平沢\t", 2)
	repeat("田井中\t", 6)
	go repeat("秋山\t", 5)
	go repeat("琴吹\t", 3)
	// repeat("田井中\t", 6)
}
