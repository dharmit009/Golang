package main

import "fmt"

func main() {
	fmt.Println("### variables.go ###")

	var myName = "Walter"
	fmt.Println("First name:", myName)

	var myLastName = "White"
	fmt.Println("Last Name:", myLastName)

	var username = "mrwhite"
	fmt.Println("Username:", username)

	var sum int
	fmt.Println("The sum is ", sum)

	part1, other := 1, 5
	fmt.Println("The part1 is ", part1, "other is", other)

	part2, other := 2, 3
	fmt.Println("The part2 is ", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("Sum: ", sum)

	var (
		lessonName = "variables"
		lessonType = "demo"
	)
	fmt.Println("lessonName", lessonName)
	fmt.Println("lessonType", lessonType)

	word1, word2, _ := "Hello", ", World", "!"
	fmt.Println(word1, word2)

}
