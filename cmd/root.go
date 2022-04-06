package cmd

import "base/pkg/infrastructure/logger"

func Execute() {
	logger := logger.NewLogger()
	logger.Info("Initiated!")
}
