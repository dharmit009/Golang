# Constant:

Expression value is known at compile time.

``` golang

const x := 1.3
const (
    y = 4
    z = "hi"
    )
```

## iota:

iota is a function which is used to create a related but
distinct constant.

**Example:** Days of the week.

Constant needs to be different but actual value is not important.

```golang

type Grades int
const (
    A Grades = iota
    B
    C
    D
    F
)

```

Each constant is assigned to unique integer. Starts at 1 and
keeps incrementing.

All you need to know about iota is they are gonna be different.
