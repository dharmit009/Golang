package main

import "fmt"

const normal, bold = "\033[0m", "\033[1m"

func main() {
	fmt.Println("What are arrays? ")
	fmt.Println("> Arrays is one of the data structure which is used to store data in contogious memory locations.")
	a := [5]int{1, 2, 3, 5}
	for x := 0; x < len(a); x++ {
		fmt.Println(a[x])
	}

}
