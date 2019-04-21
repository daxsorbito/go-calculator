package calc

// Sub method subtracts an array of int64
func Sub(args []int64) (result float32) {
	result = float32(args[0])
	for _, s := range args[1:] {
		result -= float32(s)
	}
	return
}
