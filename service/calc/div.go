package calc

// Div method divides an array of int64
func Div(args []int64) (result float64) {
	result = float64(args[0])
	for _, s := range args[1:] {
		result /= float64(s)
	}
	return
}
