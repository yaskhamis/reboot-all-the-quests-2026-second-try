package piscine

func RecursiveFactorial(nb int) int {
	// Upper Bound Check (nb > 20): Since $20!$ is the largest factorial that fits in a 64-bit integer, anything higher is guaranteed to overflow.
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
