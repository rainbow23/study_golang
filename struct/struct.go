package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

type Vertex2 struct {
	X, Y int
}

func (v Vertex2) Area() int {
	return v.X * v.Y
}

func (v *Vertex2) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex2) int {
	return v.X * v.Y
}

func main() {
	v1 := Vertex2{3, 4}
	fmt.Println(Area(v1))
	v1.Scale(10)
	fmt.Println(v1.Area())

	v := Vertex{X: 1, Y: 2}
	changeVertex(v)
	fmt.Println(v)

	v2 := Vertex{1, 2, "test"}
	changeVertex2(&v2)
	fmt.Println(&v2)
	fmt.Println(v2)
	// fmt.Println(*v2)

	/*
	 *     fmt.Println(v.X, v.Y)
	 *     v.X = 100
	 *     fmt.Println(v.X, v.Y)
	 *
	 *     v2 := Vertex{X: 1}
	 *     fmt.Println(v2)
	 *
	 *     v3 := Vertex{1, 2, "test"}
	 *     fmt.Println(v3)
	 *
	 *     v4 := Vertex{}
	 *     fmt.Println("v4", v4)
	 *
	 *     var v5 Vertex
	 *     fmt.Println(v5)
	 *
	 *     v6 := new(Vertex)
	 *     fmt.Printf("%T %v\n", v6, v6)
	 *
	 *     v7 := &Vertex{}
	 *     fmt.Printf("%T %v\n", v7, v7)
	 *
	 *     s := make([]int, 0)
	 *     // s := []int{}
	 *     fmt.Println(s)
	 */

}
