package LeetCode

import (
	"fmt"
)

func digitSum(s string, k int) string {
	if len(s) == k {
		return s
	}
	current := s
	next := ""

	for len(current) > k {

		for len(current) > 0 {
			process := ""
			if len(current) < k {
				process = current[0:]
				current = ""
			} else {
				process = current[0:k]
				current = current[k:]
			}
			sum := 0
			for _, v := range process {
				sum += int(v) - 48
			}
			next = next + fmt.Sprintf("%d", sum)

		}

		current = next
		next = ""
	}

	return current
}
