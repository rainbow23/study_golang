package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I'm at", p.Address.Number,
		p.Address.Street,
		p.Address.City,
		p.Address.State,
		p.Address.Zip)
}

type Citizen struct {
	Country string
	Person
}

func (c *Citizen) Nationality() {
	// fmt.Println(c.Person.Name, "is a citizen of", c.Country)
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

func (c *Citizen) Talk() {
	// fmt.Println("Hello, my name is", c.Person.Name, "and I'm from", c.Country)
	fmt.Println("Hello, my name is", c.Name, "and I'm from", c.Country)
}

type Human interface {
	Talk()
}

func SpeakTo(h Human) {
	h.Talk()
}

func main() {
	p := Person{
		Name: "Steve",
		Address: Address{
			Number: "13",
			Street: "Main",
			City:   "Gothan",
			State:  "NY",
			Zip:    "01313",
		},
	}
	p.Talk()
	p.Location()

	c := Citizen{}
	c.Name = "Steve"
	c.Country = "America"
	c.Talk()        // Hello, my name is Steve and I'm from America
	c.Person.Talk() // Hi, my name is Steve
	c.Nationality()

	pp := Person{Name: "Dave"}
	SpeakTo(&pp)
	cc := Citizen{Person: Person{Name: "Steve"}, Country: "America"}
	SpeakTo(&cc)
}
