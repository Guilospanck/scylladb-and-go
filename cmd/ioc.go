/* Inversion of Control */
package cmd

import (
	"base/pkg/application/interfaces"
	"base/pkg/infrastructure/database"
	"base/pkg/infrastructure/logger"
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

	return &Container{
		logger,
	}
}
