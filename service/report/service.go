package report

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Report struct holder for report objects
type Report struct {
}

func connectToDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Env>>>>", os.Getenv("MONGODB_URI"))
}

// NewReportService method to create a Report object
func NewReportService() *Report {
	connectToDB()
	return &Report{}
}
