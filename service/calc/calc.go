package calc

// Calc stuct for calc service
type Calc struct {
	Operation string
	Args      []int64
}

// NewCalcService method to create a new Calc Service
func NewCalcService(operation string, args []int64) *Calc {
	return &Calc{operation, args}
}

// Execute method to execute the operation against the arguments
func (c *Calc) Execute() (result float32) {
	switch c.Operation {
	case "add":
		result = Add(c.Args)
	case "sub":
		result = Sub(c.Args)
	case "multi":
		result = Mult(c.Args)
	case "div":
		result = Div(c.Args)
	}
	return
}
