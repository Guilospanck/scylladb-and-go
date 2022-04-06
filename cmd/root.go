package cmd

import (
	"base/pkg/infrastructure/logger"
	_ "base/pkg/infrastructure/scylladb"
)

func Execute() {
	logger := logger.NewLogger()
	logger.Info("Initiated!")
}
