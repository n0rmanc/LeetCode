package LeetCode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	maxDiff := 10000
	result := nums[0] + nums[1] + nums[2]
	l := len(nums)
	sort.Ints(nums)

	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left, right := i+1, l-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := int(math.Abs(float64(sum - target)))
			if diff < maxDiff {
				maxDiff = diff
				result = sum
			}
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return target
			}
		}

	}
	return result

}
