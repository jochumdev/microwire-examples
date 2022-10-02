package main

import (
	_ "github.com/go-micro/microwire-plugins/broker/nats/v5"
	_ "github.com/go-micro/microwire-plugins/cli/urfave/v5"
	_ "github.com/go-micro/microwire-plugins/registry/nats/v5"
	_ "github.com/go-micro/microwire-plugins/transport/nats/v5"
	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/logger"
)

func main() {
	service, err := micro.NewService(
		micro.Name("configfile"),
		micro.Usage("A POC for go-micro.dev/v5"),
		micro.Version("v0.0.1"),
		micro.NoFlags(),
		micro.ConfigFile("config"),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
