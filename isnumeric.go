package piscine

func IsNumeric(s string) bool {
	ch := []rune(s)
	for i := 0; i < len(ch); i++ {
		if !(ch[i] >= '0' && ch[i] <= '9') {
			return false
		}
	}
	return true
}
