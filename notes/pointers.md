# Pointers

Go has pointers. A pointer holds the memory address of a value. The `&` operator generates a pointer to an operand. The `*` operator denotes the pointer's underlying value. 

**Notation:** `var p *int`

``` golang

i := 32 // i is variable stored in memory. 
p = &i  // p is the pointer for i. 

fmt.Println(*p)
*p = 21

```

**Example:**

``` golang 

package main 

import "fmt"

func main(){
	i, j := 42, 2700
	p := &i

	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}

```

