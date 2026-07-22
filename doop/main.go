package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MaxInt = int64(^uint64(0) >> 1)
	MinInt = -MaxInt - 1
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	b, err2 := strconv.ParseInt(os.Args[3], 10, 64)

	if err1 != nil || err2 != nil {
		return
	}

	op := os.Args[2]

	if op == "+" {
		if (b > 0 && a > MaxInt-b) || (b < 0 && a < MinInt-b) {
			return
		}
		fmt.Println(a + b)
	} else if op == "-" {
		if (b < 0 && a > MaxInt+b) || (b > 0 && a < MinInt+b) {
			return
		}
		fmt.Println(a - b)
	} else if op == "*" {
		if a != 0 && (a*b)/a != b {
			return
		}
		fmt.Println(a * b)
	} else if op == "/" {
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		if a == MinInt && b == -1 {
			return
		}
		fmt.Println(a / b)
	} else if op == "%" {
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		if a == MinInt && b == -1 {
			return
		}
		fmt.Println(a % b)
	}
}
