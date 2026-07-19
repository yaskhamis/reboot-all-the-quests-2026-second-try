package piscine

func ConcatParams(args []string) string {
	strings := ""
	for i := 0; i < len(args); i++ {
		strings += args[i]
		if args[i] != args[len(args)-1] {
			strings += "\n"
		}
	}
	return strings
}
