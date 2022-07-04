# Printing:

You need to import `fmt` package in order to use the printing
functionalities.

``` golang
package main

import "fmt"

func main(){
    fmt.Printf("Hi")
}

```

## Format String:

Format String are which use conversion character. For Go, the
conversion character is `%`. The conversion characters further has a
character which represent the data type of the data to be
replaced.

**Here are some of the example:**

``` golang
package main

import "fmt"

func main(){
    userName := "test"
    fmt.Printf("Hello, My Name is %s", userName)
}

```
