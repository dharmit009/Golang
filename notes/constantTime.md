# Constant Time:

In constant time, the number of inputs may vary but the cost of time always
remains the same. Therefore, You can have O(n) no. of inputs but the result
would always result in O(1).

**Hint:** Accessing one element or getting constant output i.e O(1).  **Note:**
This is extremely scalable.

Here is an example of the same:

```golang

import "fmt"

func twoBoxes(boxes){
    fmt.Printf("%d", boxes[0])
    fmt.Printf("%d", boxes[1])
}

func main(){
    var boxes = []int{0,1,2,3,4,5}
    twoBoxes(boxes)
}

```

The graph form for the upper function will be of O(2) which would later be
reduced to O(1).
