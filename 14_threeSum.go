package LeetCode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	l := len(nums)
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left, right := i+1, l-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			}

		}

	}
	return result
}
