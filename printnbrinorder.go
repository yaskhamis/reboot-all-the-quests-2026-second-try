package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for i := 0; i <= 9; i++ {
		t := n
		for t > 0 {
			if t%10 == i {
				z01.PrintRune(rune('0' + i))
			}
			t /= 10
		}
	}
}
