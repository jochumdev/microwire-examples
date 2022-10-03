package main

import (
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/server"

	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/transport"

	_ "github.com/go-micro/microwire-plugins/sets/nats/v5"
)

func main() {
	// Needed to enable the logger
	loggerConfig := logger.NewConfig()
	loggerConfig.Enabled = true

	service, err := newService(
		micro.NewOptions(
			micro.Name("nocli"),
			micro.Version("v0.0.1"),
			micro.ConfigFile("config"),
		),
		client.NewConfig(),
		broker.NewConfig(),
		loggerConfig,
		registry.NewConfig(),
		server.NewConfig(),
		transport.NewConfig(),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
