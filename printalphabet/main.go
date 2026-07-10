package main

import "github.com/01-edu/z01"

// import "fmt"
func main() {
	// Note: you can't run a str or int as a rune, so you gotta make it into a rune.
	for i := 97; i <= 122; i++ {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')

	//Note: you can also write it like this:
	//for i := 'a' ; i <='z' ; i++{
	//	z01.PrintRune(i)
	//}
}

// Notes: i++ == i = i+1

// for i := 0 ; i <=10 ; i++{
// 		fmt.Print("hello \n")
// }

// you can print the numbers as well by using the for loop

// for i := 0 ; i <=10 ; i++{
// 		fmt.Println(i)
// }
