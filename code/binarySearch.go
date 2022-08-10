package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const bold = "\033[1m"
const normal = "\033[0m"

func binSearch(arr []int, target int) bool {
	panic("Functions not implementd yet")

}

func seqSearch(arr []int, target int) bool {
	panic("Functions not implementd yet")
}

func menu() {

}

func timeComplexity(options int, target int) {

	if options == 1 {
		start := time.Now()
		binSearch(target)
		end := time.Since(start)
		fmt.Println(bold+"Execution Time:", end)
	} else if options == 2 {
		start = time.Now()
		seqSearch(target)
		end = time.Since(start)
		fmt.Println(bold+"Execution Time:", end)
	} else if options == 3 {
		binSearch()
		seqSearch()
	} else {
		fmt.Println("Error [E]: Invalid Option!")
	}

}

func main() {
	fmt.Println(" ### Binary Search ### ")
	reader := bufio.NewReader(os.Stdin)
}
