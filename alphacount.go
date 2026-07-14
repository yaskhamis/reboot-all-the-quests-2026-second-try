package piscine

func AlphaCount(s string) int {
	ch := []rune(s)
	c := 0
	for i := 0; i < len(ch); i++ {
		if (ch[i] >= 'a' && ch[i] <= 'z') || (ch[i] >= 'A' && ch[i] <= 'Z') {
			c++
		}
	}
	return c
}
