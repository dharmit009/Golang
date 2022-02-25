package main

import "fmt"

type Namer struct {
	name string
}

func (n Namer) getData() string {
	fmt.Println("---------------------")
	fmt.Println("First Name: ")
	fmt.Scanf("%s", &n.name)
	fmt.Println("First Name: ", n.name)

	return "---------------------"
}

func main() {
	// Normal input ...
	var a int

	fmt.Println("vim-go")
	fmt.Println("Enter a num: ")
	fmt.Scanf("%d", &a)
	fmt.Println(a)

	// Input for structure using a method ...
	myNameStructure := Namer{}
	myNameStructure.getData()

}
