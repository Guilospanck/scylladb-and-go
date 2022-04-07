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
	session, err := database.GetConnection(logger)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	model := models.NewMutantDataTable().Table
	querybuilder := database.NewQueryBuider[entities.MutantData](model, session, logger)

	/* Insert */
	// querybuilder.Insert(&newData)

	/* Delete */
	// querybuilder.Delete(&dataToBeDeleted)

	/* Select */
	// dataToBeSearched := entities.MutantData{
	// 	FirstName: "Bob",
	// 	LastName:  "Loblaw",
	// }
	// ShowSelect[entities.MutantData](querybuilder, logger, &dataToBeSearched)

	/* Select All */
	ShowValuesSelectAll[entities.MutantData](querybuilder, logger)

	return &Container{
		logger,
	}
}
