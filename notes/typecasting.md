# Typecasting Or Type conversion:

Typecasting or Type Conversion is a process of changing a value from
one datatype to another data type. This is not always possible. In
Golang, the typecasting is pretty easy but it needs to defined
explicitly.

**Notation for Typecasting:**

``` golang
var a = dataType(value)
var b = dataType(value)

// Or

a := dataType(value)
b := dataType(value)

```

**Example:**

```golang
package main
import "fmt"

func main(){
    var x int32 = 1
    var y int16 = 2

    // this provides error as they are of different types.
    x = y

    // You will need to run the following:
    x = int32(y)
}

```
