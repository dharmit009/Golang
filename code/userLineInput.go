package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("User Line Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a line: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Line: ", name)

}
