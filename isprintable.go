package piscine

func IsPrintable(s string) bool {
	ch := []rune(s)
	for i := 0; i < len(ch); i++ {
		if !(ch[i] >= 32 && ch[i] <= 126) {
			return false
		}
	}
	return true
}
