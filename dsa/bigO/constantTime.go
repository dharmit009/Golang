package main

import "fmt"

const bold = "\033[1m"
const normal = "\033[0m"

func access(array []int) { //takes n number of inputs
	fmt.Println("Output: ", array[0]) //gives back only one
}

func main() {
	fmt.Println(bold + "What is Constant Time or O(1)?")
	var mess = normal + `
An algorithm is said to be constant time if the value of is bounded
by a value that does not depend on the size of the input. Think of it
as let there be N number of inputs but output is always pegged to one.` + bold + `

BigO Notation:` + normal + ` O(1) ` + bold + `

Example:` + normal + ` Accessing only one value from an array.` + bold + `

Code Snippet:` + normal + `

package main

import "fmt"

func main() {
	var a = []int{1, 3, 5, 6, 4, 61, 9, 92}
	fmt.Println(a[0])
}
`
	fmt.Println(mess)

	var a = []int{1, 3, 5, 6, 4, 61, 9, 92}
	fmt.Println("Input:", a)
	access(a)

}
