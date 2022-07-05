# What are arrays?

Array are of **fixed-length** series of element of a chosen type.

The elements of the array can be accessed using the subscript notation [].
Indices start at 0. Elements are initialized by zero value.

``` golang

package main

import "fmt"

func main(){
    var x [5] int
    x[0] = 2

    var y [5] int = [5] {1,2,3,4,5} //literal
    fmt.Printf(x[1]);
}
```

* Length of literal must be length of the array.
* `...` for size in array literal infers size from number of initializers.

```golang

// this sets the size of array = 3 automatically.
 x := [...]int{1,2,3,4}

 ```

## Iterating through an Array:

* Use `for` loop with `range` keyword to iterate through array.

``` golang

// i -> index & v -> value.
for i, v range x {
    fmt.Printf("Index: %d, Value: %d", i, v)
}

```

**🚩 Note:** `range` returns 2 values i.e index and value.

