# Pointers

Golang supports pointers by default. Pointers are used to hold memory
addresses. Further this memory addresses can be used to manipulate /
change / modify the underlying values which are stored in that memory
block.

| Operator | Description.                         |
|----------|--------------------------------------|
| &        | returns the address of the variable. |
| *        | returns the value of the variable.   |

**Notation:** `var p *int`

``` golang

package main

import "fmt"

func main (){
    var x int = 1
    var y int = 0

    var pointer *int = &x

    y = *pointer

    fmt.Println("Before using pointer: ");
    fmt.Println("X: ", x);
    fmt.Println("Y: ", y);
    fmt.Println("Pointer: ", pointer);

    fmt.Println("\nAfter using pointer: ");
    fmt.Println("X: ", x);
    fmt.Println("Y: ", y);
    fmt.Println("Pointer: ", pointer);

}
```

## new Function:

A alternative way to create a variable.
A new() function creates a variable and returns a pointer to the
variable.

variable is initialized to zero.

The new() function returns pointer to a variable.

``` golang
package main

import "fmt"

func main(){

    ptr := new(int)
    *ptr := 4
}

```


