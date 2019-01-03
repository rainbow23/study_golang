package main

import (
	"fmt"
	// "time"
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

// func repeat(say string, count int) {
//     for i := 0; i < count; i++ {
//         time.Sleep(1 * time.Second)
//         fmt.Println(say, fmt.Sprintf("%d回目", i))
//     }
// }

// func main() {
//     go repeat("平沢\t", 2)
//     repeat("田井中\t", 6)
//     go repeat("秋山\t", 5)
//     go repeat("琴吹\t", 3)
//     // repeat("田井中\t", 6)
// }

func main() {
	// n := make([]int, 3, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 0, 0)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	// n = append(n, 1, 2, 3, 4, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	// a := make([]int, 3)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	// a = append(a, 1)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	// a = append(a, 1, 2, 3, 4, 5)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	// b := make([]int, 0)
	// var c []int
	// fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)

	// c = make([]int, 5)
	// for i := 1; i < 6; i++ {
	//     c = append(c, i)
	//     fmt.Println(c)
	// }
	// fmt.Println(c)
	// fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	// c = make([]int, 0, 5)
	// for i := 1; i < 6; i++ {
	//     c = append(c, i)
	//     fmt.Println(c)
	// }
	// fmt.Println(c)

	// m := map[string]int{"apple": 100, "banana": 200}
	// fmt.Println(m)
	// fmt.Println(m["apple"])
	// m["banana"] = 300
	// fmt.Println(m)
	// m["new"] = 500
	// fmt.Println(m)
	// fmt.Println(m["nothing"])

	// v, ok := m["apple"]
	// fmt.Println(v, ok)

	// v2, ok2 := m["nothing"]
	// fmt.Println(v2, ok2)

	// m2 := make(map[string]int)
	// m2["pc"] = 4000
	// fmt.Println(m2)

	// // var m3 map[string]int
	// // m3["pc"] = 5000
	// // fmt.Println(m3)

	// var s []int
	// if s == nil {
	//     fmt.Println("Nil")
	// }

	// b := []byte{72, 73}
	// fmt.Println(b)
	// fmt.Println(string(b))

	// c := []byte("HI")
	// fmt.Println(c)
	// fmt.Println(string(c))

	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(xx int) {
		fmt.Println("inner func", xx)
	}(1)
}

func add(x, y int) (int, int) {
	fmt.Println("x + y")
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}
