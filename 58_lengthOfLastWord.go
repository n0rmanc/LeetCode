package LeetCode

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				break
			}
			continue
		} else {
			count++
		}
	}

	return count
}
