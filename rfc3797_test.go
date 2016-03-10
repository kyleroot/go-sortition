package rfc3797

import (
	"testing"
)

func TestEntropyOnExampleValues(t *testing.T) {
	// Values taken from Section 3.3 of RFC.
	m := map[int]int{
		20:  18,
		25:  22,
		30:  25,
		35:  28,
		40:  30,
		50:  34,
		60:  37,
		75:  40,
		100: 44,
		120: 47,
	}

	for key, value := range m {
		if Entropy(10, key) != value {
			t.Fail()
		}
	}
}
