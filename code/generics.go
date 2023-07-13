package main

import "fmt"

func length[T any](s []T) int {
	return len(s)
}

func main() {
	fmt.Println("### Generics ###")

	fmt.Println(length([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(length([]string{"hello", "world", "testing"}))

}
