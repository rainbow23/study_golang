package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./handler.go")
	if err != nil {
		log.Fatal("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}

	fmt.Println(count)
	fmt.Println(count, string(data))

	fmt.Println("----------------- error -----------------")

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error1!")
	}

	err2 := os.Chdir("test")
	if err2 != nil {
		log.Fatalln("Error2!")
	}

}
