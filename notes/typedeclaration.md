# Type Declarations:

Every variables has to have a type. So In Golang, you can make type
declarations where you actually define an alternate name for a type.

``` golang
type Celsius float64
type IDnum int

```

Once you have declared custom type you can use it as datatype while
declaring a variable.

``` golang
var temp Celsius
var rno IDnum

```

If you do not specify the type while declaring the type Golang will
infer it automatically. Though the best practice is to specify the
datatype.

``` golang
// sets x as float.
var x float = 100

// Automatically infers x as integer by interpreting the
// value of rhs.
var x = 100

```
