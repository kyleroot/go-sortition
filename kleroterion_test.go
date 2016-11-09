package kleroterion

import (
	"reflect"
	"testing"
)

func TestEntropy(t *testing.T) {
	// Values taken from Section 3.3 of the RFC.
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
		bits, _ := entropy(10, key)
		if bits != value {
			t.Fail()
		}
	}
}

func TestEntropyNGreaterThanP(t *testing.T) {
	// Trying to select more items than are in the pool.
	_, err := entropy(10, 5)
	if err == nil {
		t.Fail()
	}
}

func TestEntropyNoSelectedItems(t *testing.T) {
	// Cannot select less than one item.
	_, err := entropy(0, 5)
	if err == nil {
		t.Fail()
	}
}

func TestHash(t *testing.T) {
	// Values taken Section 6 of the RFC.
	inputs := "9319./2.5.8.10.12./9.18.26.34.41.45./"
	m := map[int][16]byte{
		0: [16]byte{0x99, 0x0D, 0xD0, 0xA5, 0x69, 0x2A, 0x02, 0x9A, 0x98, 0xB5, 0xE0, 0x1A, 0xA2, 0x8F, 0x34, 0x59},
		1: [16]byte{0x36, 0x91, 0xE5, 0x5C, 0xB6, 0x3F, 0xCC, 0x37, 0x91, 0x44, 0x30, 0xB2, 0xF7, 0x0B, 0x5E, 0xC6},
		2: [16]byte{0xFE, 0x81, 0x4E, 0xDF, 0x56, 0x4C, 0x19, 0x0A, 0xC1, 0xD2, 0x57, 0x53, 0x97, 0x99, 0x90, 0xFA},
	}

	for iter, exampleHash := range m {
		hash := hash(inputs, iter)
		if hash != exampleHash {
			t.Fail()
		}
	}
}

func TestModulo(t *testing.T) {
	// Values taken Section 6 of the RFC.
	var examples = []struct {
		dividend  [16]byte
		divisor   uint16
		remainder int16
	}{
		{[16]byte{0x99, 0x0D, 0xD0, 0xA5, 0x69, 0x2A, 0x02, 0x9A, 0x98, 0xB5, 0xE0, 0x1A, 0xA2, 0x8F, 0x34, 0x59}, 25, 16},
		{[16]byte{0x36, 0x91, 0xE5, 0x5C, 0xB6, 0x3F, 0xCC, 0x37, 0x91, 0x44, 0x30, 0xB2, 0xF7, 0x0B, 0x5E, 0xC6}, 24, 6},
		{[16]byte{0xFE, 0x81, 0x4E, 0xDF, 0x56, 0x4C, 0x19, 0x0A, 0xC1, 0xD2, 0x57, 0x53, 0x97, 0x99, 0x90, 0xFA}, 23, 1},
		{[16]byte{0x18, 0x63, 0xCC, 0xAC, 0xEB, 0x56, 0x8C, 0x31, 0xD7, 0xDD, 0xBD, 0xF1, 0xD4, 0xE9, 0x13, 0x87}, 22, 13},
	}

	for _, expl := range examples {
		rem, _ := modulo(expl.divisor, expl.dividend)
		if rem != expl.remainder {
			t.Fail()
		}
	}
}

func TestModuloByZero(t *testing.T) {
	dividend := [16]byte{0x99, 0x0D, 0xD0, 0xA5, 0x69, 0x2A, 0x02, 0x9A, 0x98, 0xB5, 0xE0, 0x1A, 0xA2, 0x8F, 0x34, 0x59}

	_, err := modulo(0, dividend)
	if err == nil {
		t.Fail()
	}
}

func TestFormat(t *testing.T) {
	// Values taken Section 6 of the RFC.
	inputs := []string{"9319", "2, 5, 12, 8, 10", "9, 18, 26, 34, 41, 45"}
	if "9319./2.5.8.10.12./9.18.26.34.41.45./" != format(inputs) {
		t.Fail()
	}
}

func TestDraw(t *testing.T) {
	names := []string{
		"John",
		"Mary",
		"Bashful",
		"Dopey",
		"Sleepy",
		"Grouchy",
		"Doc",
		"Sneazy",
		"Handsome",
		"Cassandra",
		"Pollyanna",
		"Pendragon",
		"Pandora",
		"Faith",
		"Hope",
		"Charity",
		"Lee",
		"Longsuffering",
		"Chastity",
		"Smith",
		"Pride",
		"Sloth",
		"Envy",
		"Anger",
		"Kasczynski",
	}
	inputs := []string{
		"9319",
		"2, 5, 12, 8, 10",
		"9, 18, 26, 34, 41, 45",
	}
	lotto := NewLottery(inputs, names)

	knownGoodResults := map[int]string{
		1:  "Lee",
		2:  "Doc",
		3:  "Mary",
		4:  "Charity",
		5:  "Kasczynski",
		6:  "Envy",
		7:  "Sneazy",
		8:  "Anger",
		9:  "Chastity",
		10: "Pandora",
	}
	actualResults, _ := lotto.Draw(10)
	if !reflect.DeepEqual(knownGoodResults, actualResults) {
		t.Fail()
	}
}
