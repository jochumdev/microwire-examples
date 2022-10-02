package main

import (
	"go-micro.dev/v4/logger"

	"github.com/go-micro/microwire"

	mBroker "github.com/go-micro/microwire/broker"
	mRegistry "github.com/go-micro/microwire/registry"
	mStore "github.com/go-micro/microwire/store"
	mTransport "github.com/go-micro/microwire/transport"

	_ "github.com/go-micro/microwire/plugins/transport/http"
	_ "github.com/go-micro/plugins/v4/broker/http"
	_ "github.com/go-micro/plugins/v4/registry/mdns"
)

func main() {
	service, err := newService(
		microwire.NewOptions(
			microwire.Name("nastconfigfile"),
			microwire.Usage("A POC for go-micro.dev/v5"),
			microwire.Version("v0.0.1"),
			microwire.Config("config"),
		),
		mBroker.NewConfig(),
		mRegistry.NewConfig(),
		mStore.NewConfig(),
		mTransport.NewConfig(),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
