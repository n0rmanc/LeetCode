package LeetCode

import "fmt"

func addDigits(num int) int {
	numStr := fmt.Sprintf("%d", num)
	if len(numStr) == 1 {
		return num
	}
	sum := 0
	for _, v := range numStr {
		sum += int(v) - 48
	}
	return addDigits(sum)
}
