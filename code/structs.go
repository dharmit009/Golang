package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func (p *Point) getData() {

	fmt.Println("Enter a number: ")

	fmt.Println("Enter a number: ")
}

func (p *Point) displayData() {
	fmt.Println("X: ", p.x)
	fmt.Println("Y: ", p.y)
}

func main() {
	point := Point{}
}
