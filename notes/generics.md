# Generics:

Generics are a feature in Golang that allows us to write functions and types that can work with any type.

``` golang

package main

import "fmt"

func main(){
    func length[T any](slice []T) int {
        return len(slice)
    }
    fmt.Println(length([]int{1,2,3}))
    fmt.Println(length([]string{"hello", "world"}))
}
```
