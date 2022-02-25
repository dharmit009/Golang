package main

import (
	"bufio"
	"fmt"
	"os"
)

// declaring constants for bold and normal text ...
const bold string = "\033[1m"
const normal string = "\033[0m"

// Created a Datastructure for storing Links !!
type Linker struct {
	link, lname, desc, cate string
}

func (l *Linker) getData() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(bold + "Link Name: " + normal + "")
	l.lname, _ = reader.ReadString('\n')
	fmt.Printf(bold + "Link: " + normal + "")
	l.link, _ = reader.ReadString('\n')
	fmt.Printf(bold + "Description: " + normal + "")
	l.desc, _ = reader.ReadString('\n')
	fmt.Printf(bold + "Category: " + normal + "")
	l.cate, _ = reader.ReadString('\n')
}

func (l *Linker) formatmd() {
	// markdown format generator ...
	mdf := "\n * [**" + l.lname + "**](*" + l.link + "*)\n\n**Description:**\n> " + l.desc + "\n**Category:** `" + l.cate + "`"
	fmt.Printf(mdf)
}

func main() {
	// creating Linker obj ...
	myLinker := Linker{}
	myLinker.getData()
	myLinker.formatmd()

} // main
