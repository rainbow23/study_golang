// Package main provides ...
package main

import (
	"fmt"
)

type EarthPony struct {
}

type Pony interface {
	Walk()
	Sprint()
}

//Pony interfaceを引数にとる
func Sprint(p Pony) {
	p.Walk()
	p.Walk()
	p.Walk()
}

//EarthPony構造体は Pony interfaceを実装
func (ep *EarthPony) Walk() {
	fmt.Println("歩くよ")
}

//EarthPony構造体は Pony interfaceを実装
func (ep *EarthPony) Sprint() {
	Sprint(ep)
}

type Pegasus struct {
	EarthPony
}

func (p *Pegasus) Fly() {
	fmt.Println("飛ぶよ")
}

func (p *Pegasus) Walk() {
	p.Fly()
}

func (p *Pegasus) Sprint() {
	Sprint(p)
}

func main() {
	var ep EarthPony
	fmt.Println("ep.Walk()")
	ep.Walk()
	fmt.Println("ep.Sprint()")
	ep.Sprint()
	fmt.Println("Pegasus start")

	var pg Pegasus
	pg.Walk()
	pg.Sprint()
}
