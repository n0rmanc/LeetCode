package LeetCode

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var tuples int
	n := len(nums1)
	var matrix [][]int
	matrix = append(matrix, nums1)
	matrix = append(matrix, nums2)
	matrix = append(matrix, nums3)
	matrix = append(matrix, nums4)

	for i := 0; i < 4; i++ {
		for j := 0; j < n; j++ {

		}
	}

	return tuples
}
