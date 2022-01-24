# Imports: 

**What are factored import statements?**

* Regular Import Statements: 

``` golang

import "fmt"
import "math/rand"

```

* Factored Import Statements: 

``` golang

import(
	"fmt"
	"math/rand"
)

```

## Exported Names: 

Anything which starts with a capital letter is an exported name in go for example: `Pi` of math package. 

``` golang

package main

import (
	"fmt"
	"math"
)

func main(){
	pi := math.Pi; 
	fmt.Println("The Value of Pi is ", pi)
}

```
