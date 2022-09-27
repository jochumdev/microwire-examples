//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/go-micro/microwire"
	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/registry"
	"github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"

	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/nats"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)

func myApp(opts ...mWire.Option) (*cli.App, error) {
	panic(wire.Build(
		ProvideMyServiceInitializer,
		mWire.ProvideOptions,
		mWire.ProvideMyFlags,
		mWire.ProvideApp,
	))
}

func myService(ctx *cli.Context, opts *mWire.Options) (micro.Service, error) {
	panic(wire.Build(
		ProvideMyBrokerOpts,
		broker.Provide,
		ProvideMyTransportOpts,
		transport.Provide,
		ProvideMyRegistryOpts,
		registry.Provide,
		microwire.ProvideDefaultMicroOpts,
		mWire.ProvideMicroService,
	))
}
