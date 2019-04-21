package calc

// Sub method subtracts an array of int64
func Sub(args []int64) (result float32) {
	for _, s := range args {
		result -= float32(s)
	}
	return
}
