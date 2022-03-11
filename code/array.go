package main

import "fmt"

const normal, bold = "\033[0m", "\033[1m"

func main() {
	fmt.Println("What are arrays? ")
	for x := 0; x < len(a); x++ {
		fmt.Println(a[x])
	}

}
