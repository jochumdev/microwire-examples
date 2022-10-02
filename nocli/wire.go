//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	micro "github.com/go-micro/microwire/v5"
	mBroker "github.com/go-micro/microwire/v5/broker"
	mRegistry "github.com/go-micro/microwire/v5/registry"
	mTransport "github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"
)

func newService(
	opts *micro.Options,
	brokerConfig *mBroker.Config,
	registryConfig *mRegistry.Config,
	transportConfig *mTransport.Config,
) (micro.Service, error) {
	panic(wire.Build(
		micro.ProvideConfigFile,
		micro.DiNoCliSet,
		provideService,
	))
}
