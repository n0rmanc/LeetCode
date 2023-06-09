package LeetCode

func generateParenthesis(n int) []string {
	result := _generateParenthesis("", n*2)
	// fmt.Println("len", len(result))
	return result
}

var brackets = []string{"(", ")"}

func _generateParenthesis(prefix string, n int) []string {
	result := []string{}
	if n == 0 {
		if isParenthesisValid(prefix) {
			return []string{prefix}
		} else {
			return []string{}
		}
	}
	for _, bracket := range brackets {
		result = append(result, _generateParenthesis(prefix+bracket, n-1)...)
	}
	return result
}

func isParenthesisValid(s string) bool {
	count := 0
	for _, bracket := range s {
		if bracket == '(' {
			count++
		} else if bracket == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	if count != 0 {
		return false
	}
	return true
}
