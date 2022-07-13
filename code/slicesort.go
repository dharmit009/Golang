package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = []int{1, 4, 2, 6, 8, 3}
	sort.Ints(a)
	fmt.Println(a)

}
