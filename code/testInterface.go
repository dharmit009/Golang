package main

import "fmt"

type Object interface {
	rotate()
}

type cube struct {
	l, w, h      float64
	rotation_opt []string
	rot          string
}

type tennisball struct {
	rotation_opt []string
	rot          string
}

func (c cube) area() string {
	return c.l * c.w * c.h
}

func (c cube) rotate() string {
	out = fmt.Sprinf("The cube was rotated in %s", rot)
	return out
}

func (t tennisball) rotate() string {
	if option == "LEFT" {
		out = fmt.Sprinf("That was a nice Leg Spin", rot)
	}
	out = fmt.Sprinf("That was a nice Off Spin", rot)
	return out
}

func (o Object) intForRot() string {
	o.intForRot()
}

func main() {

	mycube := cube{
		l:            3,
		w:            mycube.l,
		h:            mycube.l,
		rotation_opt: {"UP", "DOWN", "LEFT", "RIGHT"},
		rot:          rotation_opt[2],
	}

	mytennis := tennis{
		rotation_opt: []string{"LEFT", "RIGHT"},
		rot:          rotation[0],
	}

}
