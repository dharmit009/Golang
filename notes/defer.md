# Defer

`Defer` statement defers the execution of a function until the surrounding functions returns. 

**Example:**

``` golang 

package main 

import "fmt"

func main(){
	defer fmt.Println("world") 
	fmt.Println("hello") 

	fmt.Println("Testing")
	defer fmt.Println("Hehe")
	fmt.Println("Hello2")

}

```

## Stacking Defers

All the deferred functions are pushed onto `STACK`. Which then returns all the deferred calls in last-in-first-out order. 

**Example:**

``` golang 

package main

import "fmt"

func main(){
	fmt.Println("Counting"); 

	for i:=0; i<10; i++;{
		defer fmt.Println(i)
	}

	fmt.Println("done"); 
}

```

