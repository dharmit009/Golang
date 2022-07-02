# Basic Data Types:

Go has one of the most comprehensive list of data types and they are as follows:

* bool

* string

* int (Integers: Both positive and negative Integers)
	* int8
	* int16
	* int32
	* int64

* uint (Unsigned Integer: Only positive Integers)
	* uint8
	* uint16
	* uint32
	* uint64
	* uintprt: Looks like this is unsigned integer pointer.

* float (Defaults to float32)
	* float32
	* float64

* byte (alias for uint8)

* rune (alias for int32): represents Unicode code point.


* Complex Numbers
	* complex64
	* complex128

Yeah Golang even supports complex numbers (2 + 4i).

**DOUBT:** What's a Unicode Code point?

> The code points are often expressed with "U+" followed by a code in base 16 (hexadecimal). For example,
> the code point for "A" is U+0041, and for A "code point" is a slot in Unicode, representing a character.
> Depending on your choice of encoding scheme, and the number of that slot, the number.

**Example Code:**

``` golang

package	main

import (
	"fmt"
	"math/cmplx"
)

var (
	mybool bool = false
	maxInt uint64 = 1<<64-1
	equatZ complex128 = cmplx.Sqrt(-4 + 23i)
)

func main(){
	fmt.Printf("Type: %T Value: %v\n", mybool, mybool)
	fmt.Printf("Type: %T Value: %v\n", maxInt, mybool)
	fmt.Printf("Type: %T Value: %v\n", EquatZ, EquatZ)
}

```

**Zero Values:**

When no value is provided to the variables they are explicitly initialed with their zero values which depends on the data type.

Here are the zero values according to the data types:

* `0` for Numeric types.
* `False` for boolean type of values.
* `""` (empty string) for strings.









