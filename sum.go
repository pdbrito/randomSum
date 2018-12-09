// Package randomSum provides a function for generating random slices of ints that
// sum to a specified value.
package randomSum

import "math/rand"

// NIntsTotaling returns a slice of size n containing randomly generated ints
// that sum to total.
func NIntsTotaling(n, total int) (result []int) {
	currentTotal := 0

	for i := n - 1; i > 0; i-- {
		randInt := rand.Intn(total-currentTotal-i-1) + 1
		result = append(result, randInt)
		currentTotal += randInt
	}

	remainder := total - currentTotal
	result = append(result, remainder)

	return result
}
