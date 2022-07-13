# Maps:

Maps are Golangs implementation of hashtable.

* Implements of a hash table.
* Use `make()` to create a map.

* 🚩 Note:* Yes, we use `make()` to create slices and maps both.

``` golang
package main

import "fmt"

func main(){
var varName map [keyDatatype] [valueDatatype]
var idMap map [string] [int]

idMap := map [string] int {
    "Joe" : 123,
    "Jane" : 456
}
idMap["John"] = 789
idMap["Jane"] = 987

fmt.Println(idMap["Joe"])
fmt.Println(len(idMap))

delete(idMap, "Joe")

}

```

## Two Value Assignment:

``` golang

// if Joe is in idMap then p will be set to true.
// and id will be the value.
id, p := idMap["joe"]

```

## Iterate through map:

```golang

for k,v := range idMap {
    fmt.Println(k, v);
}

```
