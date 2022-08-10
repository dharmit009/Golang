# Linear Time:

When the number of operations are equal to number of elements then the algorithm
is called Linear Time Algorithm or O(n)

**Hint:** No. of Inputs == No. Of Operations

**Example:**

```golang

import "fmt"

func main(){

    var a = {"lookup", "for", "nemo"}
    for i:= 0; i < len(a); i++ {
        if a[i] == "nemo" {
            fmt.Println("Found Nemo !!")
        }
    }

}

```
