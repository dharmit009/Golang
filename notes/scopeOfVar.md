# Scope Of Variable:

The Places is code where a variable can be accessed.

**Code 1:**

``` golang

var x = 4

func f(){
    fmt.Printf("%d", x)
}
func g(){
    fmt.Printf("%d", x)
}

```

**Code 2:**

``` golang

func f(){
    var x = 4
    fmt.Printf("%d", x)
    }

func g(){
    fmt.Printf("%d", x)
}

```

In Golang, A variable scoping is done using blocks which are
represented by `{}`. This are **explicit blocks**. Golang also has **implicit Blocks** which does not use `{}`.


| Block Name                | Description                          |
|---------------------------|--------------------------------------|
| Universe Block            | All GO source code.                  |
| Package Block             | All source in a package.             |
| File Block                | All source in a file.                |
| if, for, switch           | all code inside the statement.       |
| Clauses in switch & select | Individual clauses each get a block. |




