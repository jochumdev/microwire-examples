package main

import (
	"fmt"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"gopkg.in/yaml.v2"

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

	configStore.GetCli().ConfigFile = "config"

	service, err := microwire.NewServiceWithConfigStore(
		configStore,
		microwire.Name("nastconfigfile"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.NoFlags(),
		microwire.Action(func(config microwire.ConfigStore, service micro.Service) error {
			b, err := yaml.Marshal(config)
			if err != nil {
				return err
			}
			fmt.Printf("YAML Config dump:\n\n%s\n", string(b))
			return nil
		}),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
