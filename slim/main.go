package main

import (
	_ "github.com/go-micro/microwire-plugins/broker/http/v5"
	_ "github.com/go-micro/microwire-plugins/cli/urfave/v5"
	_ "github.com/go-micro/microwire-plugins/codec/json/v5"
	_ "github.com/go-micro/microwire-plugins/registry/mdns/v5"
	_ "github.com/go-micro/microwire-plugins/transport/http/v5"
	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/logger"
)

func main() {
	service, err := micro.NewService(
		micro.Name("slim"),
		micro.Usage("A POC for go-micro.dev/v5"),
		micro.Version("v0.0.1"),
		micro.ArgPrefix(""),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
