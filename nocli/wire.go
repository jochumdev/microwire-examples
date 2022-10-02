//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/go-micro/microwire"
	mBroker "github.com/go-micro/microwire/broker"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	"github.com/google/wire"
	"go-micro.dev/v4"
)

func newService(
	opts *microwire.Options,
	brokerConfig *mBroker.Config,
	registryConfig *mRegistry.Config,
	transportConfig *mTransport.Config,
) (micro.Service, error) {
	panic(wire.Build(
		microwire.ProvideConfigFile,
		microwire.DiSet,
		provideService,
	))
}
