# For loop:

> Golang has only one looping construct and that is `for` loop. The `for` keyword is also used for `while`, `forever` aka infinite loop, and continued for loop. 

A `for` loop has three main components in it which are seperated by semicolons. The Three components are as follows:

* The init statement.
* The condition expression.
* The post statement.

The variables declared are only available within the for loop. 

**Notation for `for` loop:**

``` golang

for i:=0; i < 10; i++{
	sum += i
}

```

## Continued For loop:

The initial statement and the post statement are optional. 

``` golang

package main

import "fmt"

func main() {
	sum := 1

// Note `;` is important here. 
	for ; sum < 1000; {
			sum += sum
	}
	fmt.Println(sum)

}

```

## While Loop:

Golang uses `for` keyword for while loop too. 

**Notation:**

``` golang

for sum < 1000 {
	sum += sum
}

```

## Forever Loop:

The forever loop also uses `for` keyword. 

**Notation:**

``` golang

for{

}

```
