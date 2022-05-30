package LeetCode

func longestCommonPrefix(strs []string) string {
	sample := []byte(strs[0])
	minLen := len(sample)
	for _, str := range strs[1:] {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	current := 0
	same := []byte("")
	for i := 0; i < minLen; i++ {
		for _, str := range strs[1:] {
			if sample[i] != str[i] {
				return string(same)
			}
			current = i
		}
		same = append(same, sample[current])

	}
	return string(same)
}
