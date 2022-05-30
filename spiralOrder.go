package LeetCode

const (
	right = 0
	left  = 1
	up    = 2
	down  = 3
)

func spiralOrder(matrix [][]int) []int {
	width := len(matrix[0])
	height := len(matrix)
	result := make([]int, width*height)
	boundary := [][]int{{0, 0}, {width - 1, height - 1}}

	x, y := 0, 0
	direction := right
	for count := 0; count < width*height; count++ {
		result[count] = matrix[y][x]
		switch direction {
		case right:
			if x+1 > boundary[1][0] {
				boundary[0][0]++
				direction = down
				y++
			} else {
				x++
			}
		case left:
			if x-1 < boundary[0][1] {
				boundary[1][1]--
				direction = up
				y--
			} else {
				x--
			}
		case up:
			if y-1 < boundary[0][0] {
				boundary[0][1]++
				direction = right
				x++
			} else {
				y--
			}

		case down:
			if y+1 > boundary[1][1] {
				boundary[1][0]--
				direction = left
				x--
			} else {
				y++
			}
		}
	}
	return result
}
