# Imports: 

* What are factored import statements?

Regular Import Statements: 

``` golang

import "fmt"
import "math/rand"

```

Factored Import Statements: 

``` golang

import(
	"fmt"
	"math/rand"
)

```

# Exported Names: 

Anything which starts with a capital letter is an exported name in go for example: `pi`. 
