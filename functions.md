# Functions in Golang: 

A function in golang can be defined with zero or more arguments. The keyword used to define functions is `func`.

**Example:**

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

If the dataTypes of arguments are same you can share the type between them ... 

**Example:**

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

## Named return values: 

Named return functions are defined at the top of the function.
Named return values are generally used to document the meaning of return values. When you define functions with named return values you can use `return` keyword without any arguments this will return values you stated above. This is also known as `naked return`. 

> Note: Used naked return in smaller functions only else it can affect the readability of the code.

**Example:**

``` golang

package main

import "fmt"

func split(sum int) (x, y int){
	x = sum * 4/9
	y = sum - x
	return 
}

func main(){
	fmt.Println(split(17))

}

```

**Interesting:**

This is quite interesting Golang does not allow you to return the same variable which you passed as parameter with `naked return`. It will generate an error as duplicate argument. 

``` golang

// Invalid ...
func Sqrt(x float64) (x float64) {
	return 
}

// Invalid ...
func Sqrt(x float64) (x float64) {
	x +=1
	return 
}

// Valid ...
func Sqrt(x float64) (t float64) {
	t = x 
	return 
}

```









