package calc

// Sub method subtracts an array of int64
func Sub(args []int64) (result float64) {
	for _, s := range args {
		result -= float64(s)
	}
	return
}
