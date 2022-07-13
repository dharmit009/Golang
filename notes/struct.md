# Structs

A `struct` is a collection of fields.
It is used to Aggregate data type & group together other
objects of arbitrary type.

``` golang

package main

import "fmt"

type Cube struct {
	X int
	Y int
	Z int
}

func main(){
    c1 := new (Cube)
    c1.X = 10
	fmt.Println(Cube{1,2,3})
	fmt.Println(c1.X)
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


