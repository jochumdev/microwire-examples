package main

import (
	"os"

	"github.com/go-micro/microwire/microwire"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/nats"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)

func main() {
	app, err := myApp(
		microwire.Name("natsdefault"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.ArgPrefix(""),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
