// Package main provides ...
package main

import (
	"fmt"
)

type SomeBehavior interface {
	DoSomething() string
}

type DefaultStruct struct {
	Name string
}

func (self *DefaultStruct) DoSomething() string {
	return self.Name
}

type StructA struct {
	*DefaultStruct
}

func NewStructA(name string) *StructA {
	return &StructA{&DefaultStruct{name}}
}

type StructB struct {
	*DefaultStruct
}

func NewStructB(name string) *StructB {
	return &StructB{&DefaultStruct{name}}
}

func main() {
	var behavior SomeBehavior

	// behavior = &StructA{&DefaultStruct{"A"}}

	// 利用側は埋込構造体の存在を意識しない
	behavior = NewStructA("BB")
	fmt.Println(behavior.DoSomething())

	behavior = NewStructB("CC")
	fmt.Println(behavior.DoSomething())
}
