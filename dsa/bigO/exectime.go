package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var start = time.Now()
	var a = []string{1, 2, 3, 4}
	var end = time.Since(start)
	fmt.Println("Execution Time:", end)
}
