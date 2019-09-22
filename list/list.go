// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	// var hoge [4]int

	hoge := [4]int{1, 2, 3, 4}

	/*
	 * hoge[0] = 1
	 * hoge[1] = 1
	 * hoge[2] = 2
	 * hoge[3] = 3
	 * hoge[4] = 4
	 */

	fmt.Println(hoge[:3])
}
