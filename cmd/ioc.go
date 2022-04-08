/* Inversion of Control */
package cmd

import (
	"base/pkg/application/interfaces"
	"base/pkg/application/usecases"
	"base/pkg/domain/dtos"
	"base/pkg/infrastructure/database"
	"base/pkg/infrastructure/database/entities"
	"base/pkg/infrastructure/database/models"
	_ "base/pkg/infrastructure/environments"
	"base/pkg/infrastructure/logger"
	"base/pkg/infrastructure/repositories"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

type Container struct{}

func ShowValuesSelectAll[T any](querybuilder interfaces.IQueryBuilder[T], logger interfaces.ILogger) {
	results, err := querybuilder.SelectAll()
	if err != nil {
		logger.Error("SelectAll() error", zap.Error(err))
	}

	for _, value := range results {
		log.Printf("%+v", value)
	}
}

func ShowSelect[T any](querybuilder interfaces.IQueryBuilder[T], logger interfaces.ILogger, dataToBeSearched *T) {
	results, err := querybuilder.Select(dataToBeSearched)
	if err != nil {
		logger.Error("Select() error", zap.Error(err))
	}

	for _, value := range results {
		log.Printf("%+v", value)
	}
}

func ShowGet[T any](querybuilder interfaces.IQueryBuilder[T], logger interfaces.ILogger, dataToBeSearched *T) {
	result, err := querybuilder.Get(dataToBeSearched)
	if err != nil {
		logger.Error("Get() error", zap.Error(err))
	}

	log.Printf("%+v", result)
}

func NewContainer() *Container {
	logger := logger.NewLogger()

	/* Database connection */
	consistency := gocql.ParseConsistency(os.Getenv("SCYLLA_CONSISTENCY"))
	hosts := strings.Split(os.Getenv("SCYLLA_HOSTS"), ",")
	keyspace := "tracking"

	dbDataConnection := database.NewScyllaDBConnection(consistency, keyspace, logger, hosts...)
	session, err := database.GetConnection(dbDataConnection, logger)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	/* Query Builder */
	trackingModel := models.NewTrackingDataTable().Table
	querybuilder := database.NewQueryBuider[entities.TrackingDataEntity](trackingModel, session, logger)

	/* Repositories */
	trackingDataRepo := repositories.NewTrackingDataRepository(querybuilder, logger)

	/* Use cases */
	createTrackingDataUsecase := usecases.NewCreateTrackingDataUsecase(trackingDataRepo)

	/* Test usecase creation */
	dataToInsert := dtos.TrackingDataDTO{
		FirstName:       "Guilherme",
		LastName:        "Rodrigues",
		Timestamp:       time.Now(),
		Location:        "Brazil",
		Speed:           50.5,
		Heat:            40,
		TelepathyPowers: 10,
	}

	result, err := createTrackingDataUsecase.Perform(dataToInsert)
	if err != nil {
		logger.Error("Error on tracking data creation: ", zap.Error(err))
	}
	fmt.Printf("%+v", result)

	return &Container{}
}
