package main

import "fmt"

const normal, bold = "\033[0m", "\033[1m"

// You can have void functions in GOlang.
// NOTE: This is actually known as procedure as it does not have any
// return value.
func helloWorld() {
	fmt.Printf("Hello, " + bold + "World!" + normal)
}

// functions with return value.
func newLine() string {
	out := fmt.Sprintln("\n")
	return out
}

// functions with parameters.
func display(name string) string {
	rout := fmt.Sprintf("Hello"+bold+" %s,"+normal, name)
	rout += "\nWelcome to the World Of Go!!"
	return rout
}

func main() {
	fmt.Println("functions.go")
	helloWorld()
	fmt.Printf(newLine())
	fmt.Println(display("Gopher"))
}
