package piscine

func IsAlpha(s string) bool {
	ch := []rune(s)
	for i := 0; i < len(ch); i++ {
		if !((ch[i] >= 'A' && ch[i] <= 'Z') || (ch[i] >= 'a' && ch[i] <= 'z') || (ch[i] >= '0' && ch[i] <= '9')) {
			return false
		}
	}
	return true
}
