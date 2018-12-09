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
		total := rand.Intn(100000)
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

	fmt.Println(len(ints))
	fmt.Println(sum(ints))

	// Output:
	// 20
	// 1000
}

func sum(ints []int) int {
	sum := 0
	for _, value := range ints {
		sum += value
	}
	return sum
}
