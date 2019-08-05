// Package main provides ...
package main

import (
	"fmt"
)

type EarthPony struct {
}

type Pegasus struct {
	EarthPony
}

type Pony interface {
	Walk()
	Sprint()
}

func (ep *EarthPony) Walk() {
	fmt.Println("歩くよ")
}

func (ep *EarthPony) Sprint() {
	ep.Walk()
	ep.Walk()
	ep.Walk()
}

// func (ep *EarthPony) Sprint(self Pony) {
//     self.Walk()
//     self.Walk()
//     self.Walk()
// }

func (u *Pegasus) Fly() {
	fmt.Println("飛ぶよ")
}

//golangはEathPony構造体のWalkをoverrideしない
// https://qiita.com/shibukawa/items/16acb36e94cfe3b02aa1
func (u *Pegasus) Walk() {
	u.Fly()
}

func main() {
	// var ep EarthPony
	// ep.Walk()
	// ep.Sprint()

	var pg Pegasus
	fmt.Println("Pegasus Walk")
	pg.Walk()
	fmt.Println("Pegasus Sprint")
	pg.Sprint()
	fmt.Println("Pegasus Walk")
	pg.Walk()
	// fmt.Println('Pegasus Sprint(self *pony)')
	// pg.Sprint(pg)
}
