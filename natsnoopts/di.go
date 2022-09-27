package main

import (
	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/microwire"
	"github.com/urfave/cli/v2"
)

func ProvideMyBrokerOpts(opts *microwire.Options, c *cli.Context) *broker.BrokerOptions {
	return &broker.BrokerOptions{
		Name:      "nats",
		Addresses: "nats://localhost:4222",
	}
}

func ProvideMyServiceInitializer() microwire.InitializeServiceFunc {
	return myService
}
