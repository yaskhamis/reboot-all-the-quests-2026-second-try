package piscine

func FirstRune(s string) rune {
	ch := []rune(s)
	for range ch {
		return ch[0]
	}
	return ch[0]
}
