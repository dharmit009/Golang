# Pointers

Go has pointers. A pointer holds the memory address of a value. 
The `&` operator generates a pointer to an operand. 
The `*` operator denotes the pointer's underlying value. 

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
