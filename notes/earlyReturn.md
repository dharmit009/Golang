# Early Return:

**What is early return?**

In "Early Return", we evaluate the data using if statement. If the data
is invalid then we early return from the function.

Early returns should be used when possible for efficiency and code
clarity.

**Pros of Early Return:**

1. All Data needed for the functions is ready after error checking. 
1. Maximizes Performance.
1. Clean code.

**Example:**

``` golang

// Let us say that below line generates an error
token, err := getSession("alice"); 

// Then the below if block will early return and display the error
if err != nil {
    return 
}

// we do this to write clean code and avoid any gross errors.

```
