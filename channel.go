package main

import (
	"fmt"
)

func main() {
	// Q1. 以下の1.11をint型に変換して出力してください。
	f := 1.11

	// fmt.Println((int)f)

	i := int(f)
	fmt.Printf("%T %v\n", i, i)

	// Q2. コードを書かずに以下の出力結果を答えてください。
	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	// expect 5, 6, 2

	// Q3. 以下のコードを実行した時に
	// fmt.Printf("%T %v", m, m)
	// 以下のような出力結果となるmを作成してください。
	// map[string]int map[Mike:20 Nancy:24 Messi:30]

	// m := []map

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	// fmt.Printf("%T\n", m) //型
	// map[string]int

	// fmt.Printf("%v\n", m) //値
	// map[Nancy:24 Messi:30 Mike:20]

	fmt.Printf("%T %v\n", m, m)
}
