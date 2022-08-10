# How to measure time in Golang?

To measure time we use the built-in `time` module. To use this module first we
need to import it:

> import "time"

Once that is completed, You can ahead and use this two variables for calculating
the time taken for execution of the code.

```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
	var start = time.Now()

    // Your testing code goes here ...
	var a = []string{1, 2, 3, 4}

	var end = time.Since(start)
	fmt.Println("Execution Time:", end)
}

```

The execution time depends on various aspects such as RAM, CPU, Power, and many
other factors. Always calculate operations that a computer needs to perform
instead of calculating time as execution time can vary a lot according the
computer hardware. Always calculate the number of operations that a computer
needs to perform. 
