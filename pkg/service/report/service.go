package report

// https://medium.com/@petrousov/implementing-crud-operations-with-go-and-mongodb-cf622f2379c4
// https://github.com/eamonnmcevoy/go_rest_api/blob/master/pkg/mongo/user_service.go
// https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e

// https://gist.github.com/y0ssar1an/df2dab474520c4086926f672c52db139
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

// func connectToDB() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	mongoURI := os.Getenv("MONGODB_URI")

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		log.Fatalf("Couldn't connect to mongodb - %s", mongoURI)
// 	}
// 	err = client.Ping(ctx, readpref.Primary())
// }

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
	// Get the next result from the cursor
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
