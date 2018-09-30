package twoSum

func twoSum(nums []int, target int) []int {
	existNums := map[int]int{}
	for index, value := range nums {
		complement := target - value
		if _, ok := existNums[complement]; ok {
			return []int{existNums[complement], index}
		}
		existNums[value] = index

	}
	return []int{}
}
