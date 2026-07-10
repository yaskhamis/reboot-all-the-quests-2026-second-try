package piscine

func BasicAtoi(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		dig := int(s[i] - '0')
		res = res*10 + dig
	}
	return res
}
