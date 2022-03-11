package main

import (
	"fmt"
	"strings"
)

func main() {

	// My Decalaration Variables
	r := strings.NewReader("abcde")
	buf := make([]byte, 4)

	fmt.Println("r:=strings.NewReader('abcde')\nr:=\n", r)
	fmt.Println("buf:=make([]byte, 4)\nbuf:=\n", buf)

}
