package LeetCode

func isPalindrome(x int) bool {
	sign := 1
	numbers := []int{}
	if x <= 0 {
		x = x * -1
		sign = -1
	}
	for x > 0 {
		numbers = append(numbers, x%10)
		x = x / 10
	}
	if sign <= 0 {
		numbers = append(numbers, sign)
	}
	length := len(numbers) - 1
	if length < 0 {
		return true
	}
	for index, number := range numbers {
		if number != numbers[length-index] {
			return false
		}
	}
	return true
}
