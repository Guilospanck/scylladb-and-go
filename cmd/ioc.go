/* Inversion of Control */
package cmd

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/database"
	"base/pkg/infrastructure/database/entities"
	"base/pkg/infrastructure/database/models"
	_ "base/pkg/infrastructure/environments"
	"base/pkg/infrastructure/logger"
	"log"
	"os"
	"strings"
	"time"

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

	trackingModel := models.NewTrackingDataTable().Table
	querybuilder := database.NewQueryBuider[entities.TrackingDataEntity](trackingModel, session, logger)

	/* Insert */
	// querybuilder.Insert(&newData)

	/* Delete */
	timeLayout := "2006-01-02 15:04:05 -0700 MST"
	timestamp, err := time.Parse(timeLayout, "2017-11-11 11:05:00 +0000 UTC")
	if err != nil {
		logger.Info(err.Error())
	}

	dataToBeDeleted := entities.TrackingDataEntity{
		FirstName: "Bob",
		LastName:  "Loblaw",
		Timestamp: timestamp,
	}
	querybuilder.Delete(&dataToBeDeleted)

	/* Select */
	// dataToBeSearched := entities.TrackingDataEntity{
	// 	FirstName: "Bob",
	// 	LastName:  "Loblaw",
	// }
	// ShowSelect[entities.TrackingDataEntity](querybuilder, logger, &dataToBeSearched)

	/* Select All */
	// ShowValuesSelectAll[entities.TrackingDataEntity](querybuilder, logger)

	return &Container{
		logger,
	}
}
