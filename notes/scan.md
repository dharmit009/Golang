# Scan:

* Scan reads user input.
* Takes a pointer as an argument.
* Typed data is written to pointer.
* returns number of scanned items.


``` golang
package main

import "fmt"

func main() {

	var appleNum int
	fmt.Print("Number of Apples: ")

	fmt.Scanf("%d", &appleNum)
	fmt.Printf("%d", appleNum)
}

```
