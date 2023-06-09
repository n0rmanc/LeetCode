package LeetCode

import (
	"strings"
)

func maxVowels(s string, k int) int {
	vowels := []string{"a", "e", "i", "o", "u"}
	max := 0
	for i := range s {
		if (i + k) > len(s) {
			break
		}
		vowelCount := 0
		subString := s[i : i+k]
		for _, vowel := range vowels {
			vowelCount += strings.Count(subString, vowel)
		}
		if vowelCount > max {
			max = vowelCount
		}
	}
	return max
}
