package main

import (
	"os"

	"github.com/go-micro/microwire"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/nats"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)

func main() {
	app, err := microwire.DefaultApp(
		mWire.Name("natsdefault"),
		mWire.Usage("A POC for go-micro.dev/v5"),
		mWire.Version("v0.0.1"),
		mWire.ArgPrefix(""),

		mWire.DefaultBroker("nats"),
		mWire.DefaultTransport("nats"),
		mWire.DefaultRegistry("nats"),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
