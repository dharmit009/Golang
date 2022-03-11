package main

import "fmt"

type livingBeing interface {
	walk() string
}

type human struct {
	name string
	legs int
}

type dog struct {
	name string
	legs int
}

func (h human) walk() string {
	out := fmt.Sprintf("I am {%s} and I am walking with %d legs", h.name, h.legs)
	return out
}

func (d dog) walk() string {
	out := fmt.Sprintf("I am {%s} and I am walking with %d legs", d.name, d.legs)
	return out
}

func display(l livingBeing) {
	fmt.Println(l.walk())
}

func main() {

	h := human{name: "Steve", legs: 2}
	d := dog{}
	d.name = "Fluffy"
	d.legs = 4

	display(h)
	display(d)

}
