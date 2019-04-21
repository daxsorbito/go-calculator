package calc

// Add method adds an array of int64
func Add(args []int64) (result float32) {
	for _, s := range args {
		result += float32(s)
	}
	return
}
