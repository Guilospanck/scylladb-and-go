package cmd

import (
	_ "base/pkg/infrastructure/environments"
)

func Execute() error {
	container := NewContainer()
	println(container)

	return nil
}
