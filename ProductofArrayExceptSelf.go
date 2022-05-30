package LeetCode

func productExceptSelf(nums []int) []int {
	total := nums[0]
	zerosCount := 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == 0 {
			zerosCount++
			continue
		}
		total = total * num
	}
	result := make([]int, len(nums))
	if zerosCount >= 2 {
		return result
	}
	for i, num := range nums {
		if num == 0 {
			result[i] = total
			continue
		}
		if zerosCount >= 1 {
			result[i] = 0
			continue
		}
		result[i] = total / num
	}
	return result
}
