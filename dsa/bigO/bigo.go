package main

import "fmt"

var nemo = []string{"lookup", "nemo", "testing"}

func findnemo(array []string) {
	for i := 0; i < len(array); i++ {
		if array[i] == "nemo" {
			fmt.Printf("Found Nemo @ %d\n", i)
		}
	}
}

func main() {
	fmt.Println("Understanding Big O")
	findnemo(nemo) // linear time
}
