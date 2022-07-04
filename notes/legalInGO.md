# Legal In Go Illegal In Others ...

**Legal In GO... Illegal in Others:**

Check this snippet of code. The func `foo` says it will return a
pointer to *x* which has a value of 1. Now in the `main` function
we are calling the `foo` function and assigning the pointer to var y.

Which does not make actually any sense as the execution of func `foo`
has ended which means the variable *x* should be deallocated. But
**NO**, Golang allows the use for such cases.

``` golang

package main

import "fmt"

func foo() *int{
    x:=1
    return &x
}

func main(){
    var y *int
    y = foo()
    fmt.Printf("%d", *y)
}

```


