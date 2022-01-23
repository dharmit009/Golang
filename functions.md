# Functions in Golang: 

A function in golang can be defined with zero or more arguments. The keyword used to define functions is `func`.

Examples for functions: 

``` golang 

package main 
import "fmt"

func add(x int, y int) int{
	return x + y
}

func functionName(paraX datatypeofX, paraY datatypeOfY) returnType{
	result = paraX + paraY 
	return result
}


```

* Short Hand for same argument types. 

If the dataTypes of arguments are same you can share the type between them ... Here is an Example: 

``` golang

package main

import "fmt"

// Without short hand:
func add(x int, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

```

~I wonder what happens if i do this should probably work ...~
Yes this works as expected !!


```
package main

import "fmt"

// Without short hand:
func person(x, y, name) {
	fmt.Println("X: ", x) 
	fmt.Println("Y: ", y) 
	fmt.Println("Name: ", name) 
}

```
**DOUBT:** Can golang have functions with void return type?


## Multiple Return values: 

A function in golang can return multiple values check the example below:

``` golang

package main 

import "fmt"

func swap(a, b string) (string, string){
	return b, a
}

func main(){
	a,b := swap("hello", "world")
	fmt.Println(a,b)
}

```

