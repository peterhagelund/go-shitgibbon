# ShitGibbon

Simple utility for generating "shitgibbon" expressions.

## Installing

```bash
go get -u github.com/peterhagelund/shitgibbon
```

## Using
```go
package main

import (
	"fmt"

	"github.com/peterhagelund/shitgibbon"
)

func main() {
	fmt.Println(shitgibbon.Make())
}
```
