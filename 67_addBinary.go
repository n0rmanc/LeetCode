package LeetCode

import (
	"math"
)

func addBinary(a string, b string) string {
	maxLen := len(a)
	var result []byte = []byte(a)
	var operator []byte = []byte(b)
	diff := int(math.Abs(float64(len(a) - len(b))))
	if len(b) > len(a) {
		maxLen = len(b)
		result = []byte(b)
		operator = []byte(a)
	}

	var carry byte = 0

	for i := maxLen - 1; i >= 0; i-- {
		var oi byte = 48
		if i-diff < len(operator) && i-diff >= 0 {
			oi = operator[i-diff]
		}
		oi = oi - 48

		ri := result[i] - 48

		result[i] = ((oi)+(ri)+carry)%2 + 48
		carry = ((oi) + (ri) + carry) / 2
	}
	if carry == 1 {
		result = append([]byte{49}, result...)
	}

	return string(result)
}
