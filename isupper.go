package piscine

func IsUpper(s string) bool {
	ch := []rune(s)
	for i := 0; i < len(ch); i++ {
		if !(ch[i] >= 'A' && ch[i] <= 'Z') {
			return false
		}
	}
	return true
}
