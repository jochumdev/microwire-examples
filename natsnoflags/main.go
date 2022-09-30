package main

import (
	"go-micro.dev/v4/logger"

	"github.com/go-micro/microwire"
	_ "github.com/go-micro/microwire/plugins/cli/urfave"
	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/nats"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)

func main() {
	configStore, err := microwire.NewConfigStore()
	if err != nil {
		logger.Fatal(err)
	}

	broker := configStore.GetBroker()
	broker.Plugin = "nats"
	broker.Addresses = []string{"nats://localhost:4222"}

	registry := configStore.GetRegistry()
	registry.Plugin = "nats"
	registry.Addresses = []string{"nats://localhost:4222"}

	transport := configStore.GetTransport()
	transport.Plugin = "nats"
	transport.Addresses = []string{"nats://localhost:4222"}

	service, err := microwire.NewServiceWithConfigStore(
		configStore,
		microwire.Name("natsnoflags"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.NoFlags(),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
