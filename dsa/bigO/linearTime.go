package main

import "fmt"

const bold = "\033[1m"
const normal = "\033[0m"

func main() {

	fmt.Println(bold + "# What is Linear Time or O(n)?" + normal)
	var mess = `When number of inputs is equal to number of operations
the algorithm is said to be linear or O(n).` + bold + `

Hint:` + normal + ` No. Of Operations == No. Of Elements.` + bold + `

Code Snippet:` + normal + `

func main() {

var a = []string{"lookup", "nemo", "testing"}
	for i := 0; i < len(a); i++ {
		if a[i] == "nemo" {
			fmt.Printf("Found Nemo @ %d\n", i)
		} //if
	} //for
} //main
`

	fmt.Println(mess)

	var a = []string{"lookup", "nemo", "testing"}
	for i := 0; i < len(a); i++ {
		if a[i] == "nemo" {
			fmt.Printf("Found Nemo @ %d\n", i)
		}
	}
}
