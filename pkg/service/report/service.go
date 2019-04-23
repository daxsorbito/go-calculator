package report

import (
	"context"
	"log"
	"time"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/daxsorbito/go-calculator/models"
	"github.com/daxsorbito/go-calculator/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Report struct holder for report objects
type Report struct {
	db *mongo.Database
}

// COLLECTIONNAME collection name for the report model
const COLLECTIONNAME = "report"

// Service interface
type Service interface {
	AddReport(report model.Report) error
	GetReport(start, end time.Time) ([]*models.Report, error)
}

// NewReportService method to create a Report object
func NewReportService(db *mongo.Database) *Report {
	// connectToDB()
	return &Report{db}
}

// AddReport methond to add a report data
func (r *Report) AddReport(report *model.Report) error {
	_, err := r.db.Collection(COLLECTIONNAME).InsertOne(context.Background(), report)
	if err != nil {
		return err
	}
	return nil
}

// GetReport methond to retrieve report data
func (r *Report) GetReport(start, end strfmt.DateTime) ([]*models.Report, error) {
	cur, err := r.db.Collection(COLLECTIONNAME).Find(context.Background(),
		bson.M{
			"createddate": bson.M{
				"$gt": time.Time(start),
				"$lt": time.Time(end),
			},
		}, nil)
	if err != nil {
		log.Fatal(err)
	}
	var elements []*models.Report
	var elem *model.Report

	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		e := &models.Report{
			CreatedDate: strfmt.DateTime(elem.CreatedDate),
			Message:     elem.Message,
			Type:        elem.Type,
		}
		elements = append(elements, e)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements, nil
}
