# Comma Ok Idom:

"Comma, Ok" idom is a special case in which golang allows a variable to
be reassigned in a creation statement.

**Example:**

```
// the code below will generate an error ...

a := 1
var a = 5

```

```
// the code below will compile

a,b := 1, 2
c,b := 3, 4

```
