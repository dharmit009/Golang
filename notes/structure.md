# Structs

A `struct` is a collection of fields. 

``` golang 

package main 

import "fmt"

type Cube struct {
	X int 
	Y int
	Z int
}

func main(){
	fmt.Println(Cube{1,2,3})
}

```

## Struct Fields

``` golang

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

		func main() {
			v := Vertex{1, 2}
			v.X = 4
			fmt.Println(v.X)
		}

```
