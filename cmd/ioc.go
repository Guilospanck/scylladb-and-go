/* Inversion of Control */
package cmd

import (
	"base/pkg/application/interfaces"
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

	"github.com/gocql/gocql"
)

type Container struct{}

func ShowValuesSelectAll[T any](querybuilder interfaces.IQueryBuilder[T], logger interfaces.ILogger) {
	results, err := querybuilder.SelectAll()
	if err != nil {
		logger.Error(fmt.Sprintf("SelectAll() error %s", err.Error()))
	}

	for _, value := range results {
		log.Printf("%+v", value)
	}
}

func ShowSelect[T any](querybuilder interfaces.IQueryBuilder[T], logger interfaces.ILogger, dataToBeSearched *T) {
	results, err := querybuilder.Select(dataToBeSearched)
	if err != nil {
		logger.Error(fmt.Sprintf("Select() error %s", err.Error()))
	}

	for _, value := range results {
		log.Printf("%+v", value)
	}
}

func ShowGet[T any](querybuilder interfaces.IQueryBuilder[T], logger interfaces.ILogger, dataToBeSearched *T) {
	result, err := querybuilder.Get(dataToBeSearched)
	if err != nil {
		logger.Error(fmt.Sprintf("Get() error %s", err.Error()))
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
	// createTrackingDataUsecase := usecases.NewCreateTrackingDataUsecase(trackingDataRepo)
	// deleteTrackingDataByPrimaryKeyUsecase := usecases.NewDeleteTrackingDataByPrimaryKeyUsecase(trackingDataRepo)
	// findTrackingDataByPrimaryKey := usecases.NewFindTrackingDataByPrimaryKeyUsecase(trackingDataRepo)
	// findAllTrackingDataByPartitionKey := usecases.NewFindAllTrackingDataByPartitionKeyUsecase(trackingDataRepo)
	// findAllTrackingData := usecases.NewFindAllTrackingDataUsecase(trackingDataRepo)

	return &Container{}
}
