package main

import (
	"github.com/go-micro/microwire/v5/logger"

	micro "github.com/go-micro/microwire/v5"
	mBroker "github.com/go-micro/microwire/v5/broker"
	mRegistry "github.com/go-micro/microwire/v5/registry"
	mTransport "github.com/go-micro/microwire/v5/transport"

	_ "github.com/go-micro/microwire-plugins/broker/http/v5"
	_ "github.com/go-micro/microwire-plugins/broker/nats/v5"
	_ "github.com/go-micro/microwire-plugins/registry/mdns/v5"
	_ "github.com/go-micro/microwire-plugins/registry/nats/v5"
	_ "github.com/go-micro/microwire-plugins/transport/http/v5"
	_ "github.com/go-micro/microwire-plugins/transport/nats/v5"
)

func main() {
	service, err := newService(
		micro.NewOptions(
			micro.Name("nocli"),
			micro.Version("v0.0.1"),
			micro.ConfigFile("config"),
		),
		mBroker.NewConfig(),
		mRegistry.NewConfig(),
		mTransport.NewConfig(),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
