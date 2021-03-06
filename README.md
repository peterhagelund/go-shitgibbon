# ShitGibbon

Simple utility for generating "shitgibbon" expressions.

## Copyright and Licensing

Copyright (c) 2017-2020 Peter Hagelund

This software is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License)

See `LICENSE.txt`

## Installing

```bash
go get -u github.com/peterhagelund/go-shitgibbon
```

## Using
```go
package main

import (
	"fmt"

	"github.com/peterhagelund/shitgibbon"
)

func main() {
	generator := shitgibbon.NewGenerator()
	fmt.Println(generator.Make())
}
```
