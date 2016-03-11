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

func TestEntropyNGreaterThanP(t *testing.T) {
	// Trying to select more items than are in the pool.
	if Entropy(10, 5) != 0 {
		t.Fail()
	}
}

func TestEntropyNoSelectedItems(t *testing.T) {
	// Cannot select less than one item.
	if Entropy(0, 5) != 0 {
		t.Fail()
	}
}
