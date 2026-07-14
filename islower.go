package piscine

func IsLower(s string) bool {
	ch := []rune(s)
	for i := 0; i < len(ch); i++ {
		if !(ch[i] >= 'a' && ch[i] <= 'z') {
			return false
		}
	}
	return true
}
