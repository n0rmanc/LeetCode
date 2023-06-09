package LeetCode

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	result := [][]int{}
	for i := 0; i <= rowIndex; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, row)
	}
	return result[rowIndex]
}
