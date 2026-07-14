package piscine

func Capitalize(s string) string {
	h := ""
	c := 'a' - 'A'
	newWord := true
	for _, v := range s {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			if newWord {
				if v >= 'a' && v <= 'z' {
					v -= c
				}
				newWord = false
			} else {
				if v >= 'A' && v <= 'Z' {
					v += c
				}
			}
		} else {
			newWord = true
		}
		h += string(v)
	}
	return h
}

// I could have used the "isalpha, isupper, and islower" functions to solve this code. But I didn't know how to handle when the index get out of range.
// I wanted to make sure that onlt the first letter should be capital and the rest small letters.
