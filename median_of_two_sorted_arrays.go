package LeetCode

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	numsLeng := len(nums)
	mid := numsLeng / 2
	if numsLeng%2 == 0 {
		result = float64(nums[mid-1]+nums[mid]) / 2
	} else {
		result = float64(nums[mid])
	}
	return result
}
