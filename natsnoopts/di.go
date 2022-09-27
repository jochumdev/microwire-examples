package main

import (
	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/registry"
	"github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/urfave/cli/v2"
)

func ProvideMyBrokerOpts(opts *mWire.Options, c *cli.Context) *broker.BrokerOptions {
	return &broker.BrokerOptions{
		Name:      "nats",
		Addresses: "nats://localhost:4222",
	}
}

func ProvideMyRegistryOpts(opts *mWire.Options, c *cli.Context) *registry.RegistryOptions {
	return &registry.RegistryOptions{
		Name:      "nats",
		Addresses: "nats://localhost:4222",
	}
}

func ProvideMyTransportOpts(opts *mWire.Options, c *cli.Context) *transport.TransportOptions {
	return &transport.TransportOptions{
		Name:      "nats",
		Addresses: "nats://localhost:4222",
	}
}

func ProvideMyServiceInitializer() mWire.InitializeServiceFunc {
	return myService
}
