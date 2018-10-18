package LeetCode

import (
	"math"
	"strings"
)

func myAtoi(str string) int {

	str = strings.TrimSpace(str)
	var result, lastResult, sing int32 = 0, 0, 1
	isOverflow := false
	for index, char := range str {

		if index == 0 && char == 43 {
			sing = sing * 1
			continue
		} else if index == 0 && char == 45 {
			sing = sing * -1
			continue
		} else if char >= 48 && char <= 58 {
			lastResult = result
			result = result*10 + (int32(char) - 48)
		} else {
			break
		}
		if index != 0 {
			if result/10 != lastResult {
				isOverflow = true
				break
			}
		}
	}
	result = result * sing

	if isOverflow {
		if sing > 0 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}

	return int(result)
}
