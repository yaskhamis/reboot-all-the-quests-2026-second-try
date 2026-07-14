package piscine

func LastRune(s string) rune {
	ch := []rune(s)
	for range ch {
		return ch[len(s)-1]
	}
	return ch[len(s)-1]
}
