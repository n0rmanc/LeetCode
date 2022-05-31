package LeetCode

func isValid(s string) bool {
	symMap := map[byte]byte{
		byte(')'): byte('('),
		byte('}'): byte('{'),
		byte(']'): byte('['),
	}

	sLen := len(s)
	var comArray []byte
	for i := 0; i < sLen; i++ {
		f := s[0]
		v, ok := symMap[f]
		if !ok {
			comArray = append([]byte{f}, comArray...)
			s = s[1:]
		} else if len(comArray) > 0 {
			if v == comArray[0] {
				comArray = comArray[1:]
				s = s[1:]
			} else {
				return false
			}
		}
	}
	return len(comArray) == 0 && len(s) == 0
}
