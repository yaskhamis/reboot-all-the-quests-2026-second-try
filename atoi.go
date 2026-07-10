package piscine

func Atoi(s string) int {
	res := 0
	sign := 1
	start := 0

	if len(s) > 0 && s[0] == '-' {
		sign = -1
		start = 1
	} else if len(s) > 0 && s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		res = res*10 + int(s[i]-'0')
	}

	return res * sign
}
