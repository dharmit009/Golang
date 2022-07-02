# Variables:

One of the most important things you should know about Golang is you cannot have a `var` declared and not used. If you have one it will generate an error. This makes testing a bit of a hassle but it's also good that you will have a code that doesn't have anything which is not suppose to be there.

Just like functions you need to define your datatypes for `var`.


**Example:**

``` golang
package main

import "fmt"

var var1, var2, var3 bool

func main(){
	var i int
	fmt.Println(i, var1, var2, var3)
}

```

You can skip writing datatypes only if you initialize the var by a value.

**Example:**

``` golang

package main

import "fmt"

// normal way
var i,j int = 1,2

func main(){
	var a, b, c = true, false, "no!"
	fmt.Println(i, j, a, b, c)
}

```

## Short Hand Variable Declaration for functions:

You don't need to type `var` if you are declaring variables inside of functions, you can directly use `:=` to construct variables.

> Note: But if you are declaring variables outside the functions you cannot use `:=` construct as it is not available. You will need to use the `var` keyword.


## Constants

The `const` keyword is used to declare a constant. A `const` could be a character, string, boolean, or numeric value.

> Note: Constants cannot be declared using `:=`.

