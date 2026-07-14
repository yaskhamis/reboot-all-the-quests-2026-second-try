package piscine

func ToLower(s string) string {
	ch := []rune(s)
	for i := 0; i < len(ch); i++ {
		if ch[i] >= 'A' && ch[i] <= 'Z' {
			ch[i] = rune(ch[i]) + 32
		}
	}
	return string(ch)
}
