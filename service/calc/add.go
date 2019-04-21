package calc

// Add method adds an array of int64
func Add(args []int64) (result float64) {
	for _, s := range args {
		result += float64(s)
	}
	return
}
