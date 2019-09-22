// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	// 初期バッファサイズ
	// initialBufSize = 10000
	initialBufSize = 1
	// バッファサイズの最大値。Scannerは必要に応じこのサイズまでバッファを大きくして各行をスキャンする。
	// この値がinitialBufSize以下の場合、Scannerはバッファの拡張を一切行わず与えられた初期バッファのみを使う。
	maxBufSize = 1000000
)

func main() {
	repeat := strings.Repeat("X", 65536)
	fmt.Println(len(repeat))

	in := strings.NewReader("1st line\n" + strings.Repeat("X", 65536) + "\n3rd line\n")

	scanner := bufio.NewScanner(in)
	buf := make([]byte, initialBufSize)
	scanner.Buffer(buf, maxBufSize)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}

	/*
	 *     s := "あ"
	 *     for i := 0; i < len(s); i++ {
	 *         fmt.Printf("% x\n", s[i]) // e3 81 82
	 *     }
	 *
	 *     for _, r := range "9876543210" {
	 *         // intVal := int(r - '0') // 9, 8, 7...
	 *         // fmt.Println(intVal)
	 *         fmt.Println(r)
	 *     }
	 */

}
