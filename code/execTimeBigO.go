package main

import (
	"fmt"
	"time"
)

func testfunc(passedarray []string) {
	for x := 0; x < len(passedarray); x++ {
		fmt.Println(passedarray[x])
	}
}

func main() {
	fmt.Println("How to pass arrays to functions?")
	var myarray = []string{"Zero", "One", "Two"}
	startTime := time.Now()
	testfunc(myarray)
	fmt.Println("It took", time.Since(startTime), "to execute.")
}
