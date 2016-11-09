package kleroterion

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// MaxBitsEntropy is limited by the checksum algorithm used, in this case MD5.
const MaxBitsEntropy = 128

// Lottery represents an individual use of the sortition process.
type Lottery struct {
	pool map[int]string
	key  string
}

// NewLottery creates a new Lottery.
func NewLottery(inputs []string, names []string) *Lottery {
	var l Lottery

	l.key = format(inputs)
	l.pool = make(map[int]string)
	for i, name := range names {
		l.pool[i] = name
	}

	return &l
}

func (l *Lottery) Draw(n int) (map[int]string, error) {
	results := make(map[int]string)
	remaining := len(l.pool)

	// Initialize selected.
	selected := make([]int, remaining)
	for i := 0; i < remaining; i++ {
		// Map 0 to 1, 1 to 2, 2 to 3, ...
		selected[i] = i + 1
	}

	for i := 0; i < n; i, remaining = i+1, remaining-1 {
		hash := hash(l.key, i)
		rem, _ := modulo(uint16(remaining), hash)
		for j := 0; j < len(l.pool); j++ {
			if selected[j] > 0 {
				rem--
				if rem < 0 {
					results[i+1] = l.pool[j]
					selected[j] = 0
					break
				}
			}
		}
	}
	return results, nil
}

func format(inputs []string) string {
	var key string
	inputSuffix := "/"

	for _, input := range inputs {
		// Remove any non-float-related characters.
		reg, _ := regexp.Compile("[^0-9\\.]+")
		input := reg.ReplaceAllString(input, " ")

		// Parse floats and sort.
		var floats []float64
		values := strings.Fields(input)
		for _, value := range values {
			float, _ := strconv.ParseFloat(value, 64)
			floats = append(floats, float)
		}
		sort.Float64s(floats)

		// Format floats as specified by the RFC.
		input = ""
		for _, float := range floats {
			// Strip trailing zeros.
			trimmed := strings.TrimRight(fmt.Sprintf("%f", float), "0")
			input = input + trimmed
		}

		// Assemble the key string.
		key = key + input + inputSuffix
	}

	return key
}

func hash(key string, iter int) [16]byte {
	affix := make([]byte, 2)
	binary.BigEndian.PutUint16(affix, uint16(iter))

	// Add two-byte prefix and suffix.
	data := append(affix, []byte(key)...)
	data = append(data, affix...)

	return md5.Sum(data)
}

// Modulo operation for a very large 16-byte dividend by a comparatively small 16-bit divisor.
func modulo(divisor uint16, dividend [16]byte) (int16, error) {
	var rem uint16

	// Cannot divide by zero.
	if divisor == 0 {
		return 0, errors.New("kleroterion: cannot divide by zero")
	}

	for i := 0; i < 16; i++ {
		rem = (rem << 8) + uint16(dividend[i])
		rem %= divisor
	}

	return int16(rem), nil
}

// Determines the number of bits of entropy required to select N candidates from a pool of size P.
func entropy(n int, p int) (int, error) {
	i := 0
	result := 0.0

	// These cases represent invalid input values.
	if (n < 1) || (n >= p) {
		return 0, errors.New("kleroterion: invalid selection size")
	}

	for i = p; i > (p - n); i-- {
		result += math.Log(float64(i))
	}

	for i = n; i > 1; i-- {
		result -= math.Log(float64(i))
	}

	// Convert to the number of bits required.
	result /= math.Log(float64(2))

	return int(math.Ceil(result)), nil
}
