package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// ch := os.Args[0]
	// str := ""
	// for i := len(ch) - 1; i >= 0; i-- {
	// 	if ch[i] == '/' {
	// 		str = ch[i+1:]
	// 		break
	// 	}
	// }
	for _, v := range os.Args[0] {
		if v != '/' && v != '.' {
			z01.PrintRune(v)
		}
	}

	z01.PrintRune('\n')
}
