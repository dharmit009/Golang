package main

import "fmt"

func main() {

	var a = []int{1, 2, 3, 5, 6, 7, 8}

	for i := 0; i < len(a); i++ {
		fmt.Println("Running ...")
		if a[i] == 5 {
			fmt.Println("Found 5")
		}
	} // for loop

}
