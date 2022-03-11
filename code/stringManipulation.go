package main

import (
	"fmt"
	"strings"
)

const normal, bold = "\033[0m", "\033[1m"

func main() {

	out := "This is a test String"
	fmt.Println(strings.ToLower(out))

}
