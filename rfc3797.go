package rfc3797

import (
	"math"
)

func main() {

}

func Entropy(n int, p int) float64 {
	i := 0
	result := 0.0

	if (n < 1) || (n >= p) {
		return 0.0
	}

	for i = p; i > (p - n); i-- {
		result += math.Log(float64(i))
	}

	for i = n; i > 1; i-- {
		result -= math.Log(float64(i))
	}

	result /= 0.69315

	return result
}
