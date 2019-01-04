package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {
	// foo()
	// defer fmt.Println("world")

	// fmt.Println("hello")

	/*
	 * fmt.Println("run")
	 * defer fmt.Println(1)
	 * defer fmt.Println(2)
	 * defer fmt.Println(3)
	 * fmt.Println("success")
	 */
	file, _ := os.Open("./channel.go")
	defer file.Close()
	data := make([]byte, 500)
	file.Read(data)
	fmt.Println(string(data))
}
