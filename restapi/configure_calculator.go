// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/syllabix/swagserver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/daxsorbito/go-calculator/pkg/service"
	"github.com/daxsorbito/go-calculator/restapi/operations"
	"github.com/daxsorbito/go-calculator/restapi/operations/calc"
	"github.com/daxsorbito/go-calculator/restapi/operations/report"

	calcSrvc "github.com/daxsorbito/go-calculator/pkg/service/calc"
	reportSrvc "github.com/daxsorbito/go-calculator/pkg/service/report"
)

//go:generate swagger generate server --target ../../calculator --name Calculator --spec ../swagger.yml

func configureFlags(api *operations.CalculatorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CalculatorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	initializeEnv()
	db := initializeDB(&service.DBConfig{
		ConnectionString: os.Getenv("MONGODB_URI"),
		DBName:           os.Getenv("MONGODB_DBNAME"),
	})

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set
	api.KeyAuth = func(token string) (interface{}, error) {
		if token == "abc" {
			return token, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New(401, "incorrect api key auth")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.CalcCalcOperationHandler = calc.CalcOperationHandlerFunc(func(params calc.CalcOperationParams, principal interface{}) middleware.Responder {
		calcService := calcSrvc.NewCalcService(params.Operation, params.Args, db)
		result := calcService.Execute()

		return calc.NewCalcOperationOK().WithPayload(result)
	})

	api.ReportFindReportsHandler = report.FindReportsHandlerFunc(func(params report.FindReportsParams, principal interface{}) middleware.Responder {
		// return middleware.NotImplemented("operation report.FindReports has not yet been implemented")
		reportService := reportSrvc.NewReportService(db)
		result, _ := reportService.GetReport(*params.Start, *params.End)

		return report.NewFindReportsOK().WithPayload(result)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	swagserve := swagserver.New(swagserver.Config{
		URLPath:        "/v1/swagger",
		ShouldRender:   true,
		SwaggerSpecURL: "/swagger.json",
	})
	handleCORS := cors.Default().Handler

	return handleCORS(swagserve(handler))
}

func initializeEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func initializeDB(dbConfig *service.DBConfig) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbConfig.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	return client.Database(dbConfig.DBName)
}
