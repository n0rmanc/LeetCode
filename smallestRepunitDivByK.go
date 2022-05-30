package LeetCode

func smallestRepunitDivByK(K int) int {
	var reminder = 0
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	for lengthR := 1; lengthR <= K; lengthR++ {
		reminder = (reminder*10 + 1) % K
		if reminder == 0 {
			return lengthR
		}
	}
	return -1
}
