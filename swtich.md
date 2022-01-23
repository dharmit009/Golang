# Switch case: 

The `switch` case in Golang is similar to any other language like c, c++, or java. There are a few exception here One of them is that each and every case has `break` in it by default. This is handle by Golang automatically. Another is that you cannot use `int` or `const` in Golang `swtich` case. 

``` golang

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.\n", os)
	}
}

```
