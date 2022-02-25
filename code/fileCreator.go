package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const normal, bold = "\033[0m", "\033[1m"

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(bold + "Enter Filename: " + normal)
	filename, _ := reader.ReadString('\n')
	fmt.Printf(bold+"File Name: "+normal+"%s", filename)
	filename = strings.TrimSuffix(filename, "\n")

	creator, err := os.Create(filename)
	defer creator.Close()
	if err != nil {
		log.Fatal(err)
	}

}
