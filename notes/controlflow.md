# Control Flow:

## If statement:

`if` the condition is true then execute the consequent block else
do nothing.

``` golang
if <condition> {
    <consquent>
}

```
Go also supports else and else if statements.

``` golang

if x > 5 {
    fmt.Printf("YUP")
}
```

## For Loops:

* Iterates while a condition is true.
* May have an initialization and update operation.

There are several forms for `for` loop in Go. Here is an vanilla one:

``` golang
for <init>; <condition>; <update> {
    <statements>
}

```

**Other forms:**

``` golang

// while loop styled
i = 0
for i < 10 {
    fmt.Printf("hi")
    i++
}

//forever loop
for {
    fmt.Printf("hi")
}

```

## Switch / Case:

Switch is a multi-way if statement. Your switch may contain a tag which
is variable to be checked.

* Tags are compared to a constant defined in each case.
* Case gets executed once the tag gets matched up to the constant
defined for each case.

```golang

switch x {
case 1:
    fmt.Printf("case1");

case 2:
    fmt.Printf("case2");

default:
    fmt.Printf("nocase");

```

The switch case is automatically `break` you don't need to do it
manually. The `default` case is completely an optional case.

## Tagless Switch:

Switch may not contain a tag. In this case contains a boolean expression
which is further evaluated. First true case gets executed.

``` golang

switch {

case x > 1:
    fmt.Printf("case 1");
case x < -1:
    fmt.Printf("case 2");
default:
    fmt.Printf("No case");

```

## Break and Continue:

``` golang

for i < 10 {
    i ++
    if i == 5 {break}
    fmt.Printf("hi");
}

for i < 10 {
    i ++
    if i == 5 { continue }
    fmt.Printf("hi");
}

```



