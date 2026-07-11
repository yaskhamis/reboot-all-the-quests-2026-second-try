package piscine

func IterativeFactorial(nb int) int {
	n := 1
	// Upper Bound Check (nb > 20): Since $20!$ is the largest factorial that fits in a 64-bit integer, anything higher is guaranteed to overflow.
	if nb < 0 || nb > 20 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		n *= i
	}
	return n
}
