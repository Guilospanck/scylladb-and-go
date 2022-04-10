package cmd

import (
	_ "base/pkg/infrastructure/environments"
)

func Execute() error {
	container := NewContainer()

	container.httpServer.Setup()

	container.trackingPresenter.Register(container.httpServer)

	if err := container.httpServer.Run(); err != nil {
		return err
	}

	return nil
}
