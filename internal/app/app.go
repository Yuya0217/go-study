package app

import (
	"go-layered-architecture-sample/internal/config"
	"go-layered-architecture-sample/internal/presentation/rest"
	"go-layered-architecture-sample/pkg/logger"
)

func Run(config *config.Config, logger logger.Logger) error {
	restServer, err := rest.NewServer(config, logger)
	if err != nil {
		return err
	}

	restServer.Start()

	return nil
}
