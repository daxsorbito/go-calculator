package calc

import (
	"fmt"
	"log"
	"time"

	reportSrvc "github.com/daxsorbito/go-calculator/pkg/service/report"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/daxsorbito/go-calculator/pkg/model"
)

// Calc stuct for calc service
type Calc struct {
	Operation string
	Args      []int64
	db        *mongo.Database
}

// NewCalcService method to create a new Calc Service
func NewCalcService(operation string, args []int64, db *mongo.Database) *Calc {
	return &Calc{operation, args, db}
}

// Execute method to execute the operation against the arguments
func (c *Calc) Execute() (result float32) {
	defer func() {
		msg := fmt.Sprintf("%s: %v", c.Operation, c.Args)
		reportType := "operation"
		if r := recover(); r != nil {
			msg = fmt.Sprintf("%v", r)
			reportType = "error"
		}
		m := &model.Report{
			Type:        reportType,
			Message:     msg,
			CreatedDate: time.Now(),
		}
		service := reportSrvc.NewReportService(c.db)
		err := service.AddReport(m)
		if err != nil {
			log.Fatal(err)
		}
	}()
	switch c.Operation {
	case "add":
		result = Add(c.Args)
	case "sub":
		result = Sub(c.Args)
	case "multi":
		result = Mult(c.Args)
	case "div":
		result = Div(c.Args)
	case "sqrt":
		result = Sqrt(c.Args[0])
	case "cbrt":
		result = Cbrt(c.Args[0])
	case "pow":
		result = Pow(c.Args[0], c.Args[1])
	case "fac":
		result = Fac(c.Args[0])
	}
	return
}
