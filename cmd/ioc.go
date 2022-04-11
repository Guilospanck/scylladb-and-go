/* Inversion of Control */
package cmd

import (
	"base/pkg/application/usecases"
	"base/pkg/infrastructure/database"
	"base/pkg/infrastructure/database/entities"
	"base/pkg/infrastructure/database/models"
	_ "base/pkg/infrastructure/environments"
	httpserver "base/pkg/infrastructure/http_server"
	"base/pkg/infrastructure/logger"
	"base/pkg/infrastructure/repositories"
	"base/pkg/interfaces/http/handlers"
	"base/pkg/interfaces/http/presenters"
	"os"
	"strings"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Container struct {
	httpServer        httpserver.IHTTPServer
	trackingPresenter presenters.IRoutes
	dbSession         *gocqlx.Session
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

	/* Query Builder */
	trackingModel := models.NewTrackingDataTable().Table
	querybuilder := database.NewQueryBuider[entities.TrackingDataEntity](trackingModel, session, logger)

	/* Repositories */
	trackingDataRepo := repositories.NewTrackingDataRepository(querybuilder, logger)

	/* Use cases */
	createTrackingDataUsecase := usecases.NewCreateTrackingDataUsecase(trackingDataRepo)
	deleteTrackingDataByPrimaryKeyUsecase := usecases.NewDeleteTrackingDataByPrimaryKeyUsecase(trackingDataRepo)
	// findTrackingDataByPrimaryKey := usecases.NewFindTrackingDataByPrimaryKeyUsecase(trackingDataRepo)
	// findAllTrackingDataByPartitionKey := usecases.NewFindAllTrackingDataByPartitionKeyUsecase(trackingDataRepo)
	// findAllTrackingData := usecases.NewFindAllTrackingDataUsecase(trackingDataRepo)

	/* HTTP server */
	httpServer := httpserver.NewHTTPServer(logger)

	/* Handlers */
	trackingHandler := handlers.NewTrackingDataHandler(logger, createTrackingDataUsecase,
		deleteTrackingDataByPrimaryKeyUsecase)

	/* Routes (Presenters) */
	trackingPresenter := presenters.NewTrackingDataPresenters(logger, trackingHandler)

	return &Container{
		httpServer:        httpServer,
		trackingPresenter: trackingPresenter,
		dbSession:         session,
	}
}
