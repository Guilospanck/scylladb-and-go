/* Inversion of Control */
package cmd

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/database"
	"base/pkg/infrastructure/logger"
	"log"

	"github.com/scylladb/gocqlx/v2/table"
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
	logger.Info("DB Session created!")

	mutantDataMetadata := table.Metadata{
		Name:    "mutant_data",
		Columns: []string{"first_name", "last_name", "address", "picture_location"},
		PartKey: []string{"first_name", "last_name"},
	}

	mutantDataTable := table.New(mutantDataMetadata)

	type MutantData struct {
		FirstName       string
		LastName        string
		Address         string
		PictureLocation string
	}

	selectStatement, statementNames := mutantDataTable.SelectAll()
	selectQuery := session.Query(selectStatement, statementNames)

	var results []MutantData
	err = selectQuery.SelectRelease(&results)
	if err != nil {
		logger.Error("Select() error", zap.Error(err))
	}

	for _, value := range results {
		log.Printf("%+v", value)
	}

	return &Container{
		logger,
	}
}
