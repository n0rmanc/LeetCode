package LeetCode

import (
	// "fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := _fourSum([]int{}, nums, target)
	return result
}

func _fourSum(picked []int, nums []int, target int) [][]int {
	if len(picked) >= 4 {

		sum := 0
		for _, v := range picked {
			sum += v
		}

		if sum == target {
			return [][]int{append([]int{}, picked...)}
		}
		return [][]int{}
	}

	if len(nums) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0, len(nums))

	for i, num := range nums {
		newPicked := append(picked, num)

		r := _fourSum(newPicked, nums[i+1:], target)

		if len(r) == 0 {
			continue
		}

		isInclude := false
		for _, rr := range result {
			if testEq(r[0], rr) {
				isInclude = true
				break
			}

		}
		if isInclude {
			continue
		}
		result = append(result, r...)

	}

	return result
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
