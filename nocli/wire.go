//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"
)

func newService(
	opts *micro.Options,
	brokerConfig *broker.Config,
	registryConfig *registry.Config,
	transportConfig *transport.Config,
) (micro.Service, error) {
	panic(wire.Build(
		micro.ProvideConfigFile,
		micro.DiNoCliSet,
		provideService,
	))
}
