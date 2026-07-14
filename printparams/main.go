package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ch := os.Args
	for i := 1; i < len(ch); i++ {
		for _, v := range ch[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
