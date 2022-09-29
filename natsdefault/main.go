package main

import (
	"github.com/go-micro/microwire"
	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/registry"
	"github.com/go-micro/microwire/transport"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/microwire/plugins/cli/urfave"
	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/nats"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)

func main() {
	service, err := microwire.NewService(
		microwire.Name("natsdefault"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.ArgPrefix(""),

		microwire.Component(broker.ComponentName, "nats"),
		microwire.Component(registry.ComponentName, "nats"),
		microwire.Component(transport.ComponentName, "nats"),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
