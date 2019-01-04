package main

import (
	"fmt"
)

func main() {
	fmt.Println("slice")
	var s1 []int
	var s2 = []int{1, 2, 3, 4}
	s3 := []int{1, 2, 3, 4}

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	s3s := s3[0:2]
	fmt.Println(s3s, len(s3s), cap(s3s))

	s3sc := s3[0:2:3]
	fmt.Println(s3sc, len(s3sc), cap(s3sc))

	fmt.Println("map")
	var m1 map[int]string

	var m2 = map[string]int{
		"dora": 0,
		"nobi": 1,
	}

	m3 := map[string]int{
		"shizu": 0,
		"jai":   1,
		"sune":  2,
	}

	fmt.Println(m1, len(m1))
	fmt.Println(m2, len(m2))
	fmt.Println(m3, len(m3))
}
