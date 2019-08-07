// Package main provides ...
package main

import (
	"fmt"
)

type Strategy interface {
	DoSomething()
}

type ConcreteStrategy struct {
}

func (self *ConcreteStrategy) DoSomething() {
	fmt.Println("ConcreteStrategy DoSomething")
}

type ConcreteStrategy2 struct {
}

func (self *ConcreteStrategy2) DoSomething() {
	fmt.Println("ConcreteStrategy2 DoSomething")
}

type StructA struct {
	strategy Strategy
}

func (self *StructA) Operation() {
	self.strategy.DoSomething()
}

func main() {
	structA := &StructA{&ConcreteStrategy{}}
	structA.Operation()
	structA.strategy = &ConcreteStrategy2{}
	structA.Operation()

}
