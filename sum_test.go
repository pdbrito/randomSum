package randomSum_test

import (
	"fmt"
	. "github.com/pdbrito/randomSum"
	"math/rand"
	"testing"
	"testing/quick"
	"time"
)

func TestNIntsTotaling(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	f := func(n int) bool {
		total := rand.Intn(100000) + 1
		size := rand.Intn(total-1) + 1

		ints := NIntsTotaling(size, total)

		return len(ints) == size && sum(ints) == total
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func ExampleNIntsTotaling() {
	ints := NIntsTotaling(20, 1000)

	fmt.Printf("ints: %v\n", ints)
	fmt.Printf("The length of ints is %v\n", len(ints))
	fmt.Printf("The sum of ints is %v\n", sum(ints))

	// Output:
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
