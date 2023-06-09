package LeetCode

var (
	numberToAlphabet = map[string][]string{
		"1": {},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
		"0": {},
	}
)

func letterCombinations(digits string) []string {

	result := tmp("", digits)

	return result
}

func tmp(prefix string, digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		if len(prefix) == 0 {
			return result
		}
		return []string{prefix}
	}
	alphabets := numberToAlphabet[string(digits[0])]
	for _, alphabet := range alphabets {
		result = append(result, tmp(prefix+alphabet, digits[1:])...)
	}
	return result
}
