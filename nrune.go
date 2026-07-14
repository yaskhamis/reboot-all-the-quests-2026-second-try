package piscine

func NRune(s string, n int) rune {
	ch := []rune(s)
	if n <= 0 || len(ch) < n {
		return 0
	}
	return ch[n-1]
}
