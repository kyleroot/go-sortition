package rfc3797

import (
	"math"
	"os"
)

func main() {
	os.Exit(0)
}

func usage() {
	os.Exit(2)
}

func Entropy(n int, p int) int {
	i := 0
	result := 0.0

	// These cases represent invalid input values.
	if (n < 1) || (n >= p) {
		return 0
	}

	for i = p; i > (p - n); i-- {
		result += math.Log(float64(i))
	}

	for i = n; i > 1; i-- {
		result -= math.Log(float64(i))
	}

	// Convert to the number of bits reqr'd.
	result /= math.Log(float64(2))

	return int(math.Ceil(result))
}
