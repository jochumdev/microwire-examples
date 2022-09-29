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

	configStore.GetBroker().Plugin = "nats"
	configStore.GetRegistry().Plugin = "nats"
	configStore.GetTransport().Plugin = "nats"

	service, err := microwire.NewServiceWithConfigStore(
		configStore,
		microwire.Name("natscompileconfig"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
