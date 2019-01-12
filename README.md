# randomSum
[![Build Status](https://travis-ci.com/pdbrito/randomSum.png?branch=master)](https://travis-ci.com/pdbrito/randomSum) [![GoDoc](https://godoc.org/github.com/pdbrito/randomSum?status.svg)](https://godoc.org/github.com/pdbrito/randomSum) [![Go Report Card](https://goreportcard.com/badge/github.com/pdbrito/randomSum)](https://goreportcard.com/report/github.com/pdbrito/randomSum)

Package randomSum  provides a function for generating random slices of ints that sum to a specified value.

### Example

```go
package main

import (
	"fmt"
	"github.com/pdbrito/randomSum"
)

func main() {
	ints := randomSum.NIntsTotaling(20, 1000)
	fmt.Printf("ints: %v\n", ints)
	fmt.Printf("The length of ints is %v\n", len(ints))
	fmt.Printf("The sum of ints is %v\n", sum(ints))

	// ints: [42 862 12 56 2 5 4 1 1 1 3 2 1 1 1 1 1 1 1 2]
	// The length of ints is 20
	// The sum of ints is 1000
}

func sum(ints []int) int {
	sum := 0
	for _, value := range ints {
		sum += value
	}
	return sum
}
```

### Seeding for random values
randomSum relies on the [rand package](https://golang.org/pkg/math/rand) to generate values for the ints that 
it returns.

By default this is deterministic and can only be changed by [setting
a seed](https://golang.org/pkg/math/rand/#Seed) **before** calling the NIntsTotalling function.

E.g. `rand.Seed(time.Now().UnixNano())`