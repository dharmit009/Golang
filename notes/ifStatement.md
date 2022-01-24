# If Statement: 

The keyword `if` is used to declare the statement. It is not necessary to sound the expression by parentheses `( )` but `{ }` are required just like `for` loop. 

**Notation:**

``` golang 

if condition {
	//statments
}

x:=4
if x<4 {
	fmt.Println("Value: %v", x)
}

```

## If statement with a short statement 

``` golang 

x:=4
if y:=4; x==4 {
	fmt.Println("The value of x is 4")
	fmt.Println("%v", y)
}

```

## If Else statement

The variables declared `if` short hand are also available within `else` blocks so make sure to use them wisely. 

**Notation:**

``` golang

if x:=0; x == 1 {
	fmt.Print("%v", x)
	fmt.Println("The value of x is 1")
}
else{
	fmt.Print("%v", x)
	fmt.Println("The value of x is 0")

}

```













