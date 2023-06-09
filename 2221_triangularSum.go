package LeetCode

func triangularSum(nums []int) int {
	i := 0
	l := len(nums)
	for l > 1 {
		for i+1 < l {
			nums[i] = (nums[i] + nums[i+1]) % 10
			i++
		}
		l--
		i = 0
	}
	return nums[0]
}
