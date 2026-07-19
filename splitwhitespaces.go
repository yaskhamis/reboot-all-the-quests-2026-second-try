package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	out := []string{}

	for _, v := range s {
		if v == ' ' || v == '\t' || v == '\n' {
			if str != "" {
				out = append(out, str)
				str = ""
			}
		} else {
			str += string(v)
		}
	}
	if str != "" {
		out = append(out, str)
	}
	return out
}
