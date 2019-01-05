package main

import "fmt"

func searchMin(first int, second int) int {
	if first > second {
		return second
	}
	return first
}
func main() {
	// Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	// l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4, 3, 50, 56, 22, 1}
	// fmt.Println("len ", len(l))
	// fmt.Println("l[i]", l[8])
	var min int
	for i := 0; i < len(l); i++ {
		if i == len(l)-1 {
			// fmt.Printf("break %T %v %d\n", len(l), len(l), len(l))
			break
		}

		// fmt.Printf("first:%v second:%v\n", l[i], l[j])
		// fmt.Printf("first:%v second:%v\n", l[i], l[i+1])
		min = searchMin(l[i], l[i+1])
	}
	fmt.Println(min)

	// Q2. 以下の果物の価格の合計を出力するコードを書いてください。
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	sum := calc_sum(m)
	fmt.Println("sum ", sum)
}

func calc_sum(fruits map[string]int) int {
	var sum int
	for _, v := range fruits {
		sum += v
	}
	return sum
}
