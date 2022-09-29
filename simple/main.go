package main

import (
	"github.com/go-micro/microwire"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/microwire/plugins/cli/urfave"
	_ "github.com/go-micro/microwire/plugins/transport/http"
	_ "github.com/go-micro/plugins/v4/broker/http"
	_ "github.com/go-micro/plugins/v4/registry/mdns"
)

func main() {
	service, err := microwire.NewService(
		microwire.Name("livecyclehooks"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.ArgPrefix(""),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
