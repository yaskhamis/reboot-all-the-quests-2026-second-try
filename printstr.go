package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	arr := []rune(s)
	for _, char := range arr {
		z01.PrintRune(char)
	}
}
