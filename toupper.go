package piscine

func ToUpper(s string) string {
	str := []rune(s)

	c := 'a' - 'A'
	h := ""

	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			if str[i] >= 'a' && str[i] <= 'z' {
				h += string(str[i] - c)
			} else if str[i] >= 'A' && str[i] <= 'Z' {
				h += string(str[i])
			}
		} else {
			h += string(str[i])
		}
	}
	return h
}
