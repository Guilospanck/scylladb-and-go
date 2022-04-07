package cmd

import (
	_ "base/pkg/infrastructure/environments"
)

func Execute() error {
	NewContainer()

	return nil
}
