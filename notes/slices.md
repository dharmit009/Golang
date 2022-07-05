# Slices:

A slice is a window on an underlying array.

**Advantages Of Slices:**

* Variable size: You can increase the size of the array.
* Pointer: Indicates the start of the slice.
* Length: Is the number of elements(elts) in the slice.
* Capacity: From the start of the slice to the end of the array.

```golang

arr:= [...] string {"a", "b", "c", "d", "e", "f"}

s1 := arr [1:3]
s2 := arr [2:5]

```

## Length and Capacity.

* len()
* cap()


```golang
package main

import "fmt"

func main(){
    a := [3] string {"a", "b", "c"}
    slice1 := a1[0:1]
    fmt.Printf(len(slice1), cap(slice1));
}

```

* Writing to slice changes underlying array.
* Overlapping slices refer to the same array elements.

**Slice Literals:**

* They can be used to initialize a slice.
* Creates the underlying array & creates a slice to reference it.
* Slice points to the start of the array, length is capacity.

``` golang

slice1 := [] int {1,2,3} // this makes a slice.

```

## Variable Slice:

* `make()`: It can be used to make a slice and you don't need to initialize any value.

**2 Arg Version:**
Specify type and length/capacity.

Initialize to zero, Length = Capacity

slice1 := make([]int, 10)

**3 Arg Version:**
Specify length and capacity seperately.

slice1 = make([]int, 10, 12);

## Append

To add elements to the end of slice we use `append()` function. This also results in increase in size of slice.
If you have reached the capacity of the array you can still append it will enlarge the array as per requirement.

* Inserts into underlying array.

``` golang

sli = make([]int, 0, 3)
sli = append(sli, 100)

```
