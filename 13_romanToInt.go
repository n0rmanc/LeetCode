package LeetCode

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var v, lv, cv int

	for i := len(s) - 1; i >= 0; i-- {
		cv = romanMap[string(s[i])]
		if cv >= lv {
			v += cv
		} else {
			v -= cv
		}
		lv = cv
	}
	return v
}
