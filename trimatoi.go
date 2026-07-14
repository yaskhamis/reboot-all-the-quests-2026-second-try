package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundDigit := false
	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
			foundDigit = true
		} else if !foundDigit && r == '-' {
			sign = -1
		} else if !foundDigit && r == '+' {
			sign = 1
		}
	}
	return result * sign
}
