/* Inversion of Control */
package cmd

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/database"
	"base/pkg/infrastructure/database/entities"
	"base/pkg/infrastructure/database/models"
	"base/pkg/infrastructure/logger"
	"log"
	"os"
	"strings"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

type Container struct {
	logger interfaces.ILogger
}

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

func NewContainer() *Container {
	logger := logger.NewLogger()

	consistency := gocql.ParseConsistency(os.Getenv("SCYLLA_CONSISTENCY"))
	hosts := strings.Split(os.Getenv("SCYLLA_HOSTS"), ",")
	keyspace := "catalog"

	mutantDataConnection := database.NewScyllaDBConnection(consistency, keyspace, logger, hosts...)

	keyspace = "tracking"
	trackingDataConnection := database.NewScyllaDBConnection(consistency, keyspace, logger, hosts...)

	/* Database connection */
	mutantSession, err := database.GetConnection(mutantDataConnection, logger)
	if err != nil {
		panic(err)
	}
	defer mutantSession.Close()

	trackingSession, err := database.GetConnection(trackingDataConnection, logger)
	if err != nil {
		panic(err)
	}
	defer trackingSession.Close()

	mutantModel := models.NewMutantDataTable().Table
	mutantQuerybuilder := database.NewQueryBuider[entities.MutantDataEntity](mutantModel, mutantSession, logger)

	trackingModel := models.NewTrackingDataTable().Table
	trackingQuerybuilder := database.NewQueryBuider[entities.TrackingDataEntity](trackingModel, trackingSession, logger)

	/* Insert */
	// querybuilder.Insert(&newData)

	/* Delete */
	// querybuilder.Delete(&dataToBeDeleted)

	/* Select */
	// dataToBeSearched := entities.MutantDataEntity{
	// 	FirstName: "Bob",
	// 	LastName:  "Loblaw",
	// }
	// ShowSelect[entities.MutantDataEntity](querybuilder, logger, &dataToBeSearched)

	/* Select All */
	ShowValuesSelectAll[entities.MutantDataEntity](mutantQuerybuilder, logger)
	ShowValuesSelectAll[entities.TrackingDataEntity](trackingQuerybuilder, logger)

	return &Container{
		logger,
	}
}
