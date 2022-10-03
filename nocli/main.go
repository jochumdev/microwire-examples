package main

import (
	"github.com/go-micro/microwire/v5/logger"

	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/transport"

	_ "github.com/go-micro/microwire-plugins/sets/nats/v5"
)

func main() {
	service, err := newService(
		micro.NewOptions(
			micro.Name("nocli"),
			micro.Version("v0.0.1"),
			micro.ConfigFile("config"),
		),
		broker.NewConfig(),
		registry.NewConfig(),
		transport.NewConfig(),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
