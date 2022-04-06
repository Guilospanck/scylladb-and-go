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

	_, err := database.GetConnection(logger)
	if err != nil {
		panic(err)
	}
	logger.Info("DB Session created!")

	return &Container{
		logger,
	}
}
