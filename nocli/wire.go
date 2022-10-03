//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"
)

func newService(
	opts *micro.Options,
	clientConfig *client.Config,
	brokerConfig *broker.Config,
	loggerConfig *logger.Config,
	registryConfig *registry.Config,
	serverConfig *server.Config,
	transportConfig *transport.Config,
) (micro.Service, error) {
	panic(wire.Build(
		micro.ProvideConfigFile,
		micro.DiNoCliSet,
		provideService,
	))
}
