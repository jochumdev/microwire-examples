//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/microwire"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"

	_ "github.com/go-micro/plugins/v4/broker/nats"
	_ "github.com/go-micro/plugins/v4/registry/nats"
	_ "github.com/go-micro/plugins/v4/transport/nats"
)

func myApp(opts ...microwire.Option) (*cli.App, error) {
	panic(wire.Build(
		ProvideMyServiceInitializer,
		microwire.ProvideOptions,
		microwire.ProvideMyFlags,
		microwire.ProvideApp,
	))
}

func myService(ctx *cli.Context, opts *microwire.Options) (micro.Service, error) {
	panic(wire.Build(
		ProvideMyBrokerOpts,
		broker.ProvideBroker,
		microwire.ProvideDefaultMicroOpts,
		microwire.ProvideMicroService,
	))
}
