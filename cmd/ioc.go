/* Inversion of Control */
package cmd

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/database"
	"base/pkg/infrastructure/database/entities"
	"base/pkg/infrastructure/database/models"
	"base/pkg/infrastructure/logger"
	"log"

	"go.uber.org/zap"
)

type Container struct {
	logger interfaces.ILogger
}

func NewContainer() *Container {
	logger := logger.NewLogger()

	/* Database connection */
	session, err := database.GetConnection(logger)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	model := models.NewMutantDataTable().Table
	querybuilder := database.NewQueryBuider[entities.MutantData](model, session, logger)

	results, err := querybuilder.SelectAll()
	if err != nil {
		logger.Error("SelectAll() error", zap.Error(err))
	}

	for _, value := range results {
		log.Printf("%+v", value)
	}

	return &Container{
		logger,
	}
}
