// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/config/configdi"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/transport"
)

import (
	_ "github.com/go-micro/microwire-plugins/sets/nats/v5"
)

// Injectors from wire.go:

func newService(opts *micro.Options, clientConfig *client.Config, brokerConfig *broker.Config, loggerConfig *logger.Config, registryConfig *registry.Config, serverConfig *server.Config, transportConfig *transport.Config) (micro.Service, error) {
	diConfig, err := micro.ProvideConfigFile(opts)
	if err != nil {
		return nil, err
	}
	config, err := configdi.ProvideConfigor(diConfig)
	if err != nil {
		return nil, err
	}
	brokerDiConfig, err := broker.ProvideConfigNoFlags(brokerConfig, config)
	if err != nil {
		return nil, err
	}
	loggerDiConfig, err := logger.ProvideConfigNoFlags(loggerConfig, config)
	if err != nil {
		return nil, err
	}
	loggerLogger, err := logger.Provide(loggerDiConfig, loggerConfig)
	if err != nil {
		return nil, err
	}
	registryDiConfig, err := registry.ProvideConfigNoFlags(registryConfig, config)
	if err != nil {
		return nil, err
	}
	registryRegistry, err := registry.Provide(registryDiConfig, loggerLogger, registryConfig)
	if err != nil {
		return nil, err
	}
	brokerBroker, err := broker.Provide(brokerDiConfig, loggerLogger, registryRegistry, brokerConfig)
	if err != nil {
		return nil, err
	}
	clientDiConfig, err := client.ProvideConfigNoFlags(clientConfig, config)
	if err != nil {
		return nil, err
	}
	transportDiConfig, err := transport.ProvideConfigNoFlags(transportConfig, config)
	if err != nil {
		return nil, err
	}
	transportTransport, err := transport.Provide(transportDiConfig, loggerLogger, transportConfig)
	if err != nil {
		return nil, err
	}
	clientClient, err := client.Provide(clientDiConfig, brokerBroker, loggerLogger, registryRegistry, transportTransport, clientConfig)
	if err != nil {
		return nil, err
	}
	serverDiConfig, err := server.ProvideConfigNoFlags(serverConfig, config)
	if err != nil {
		return nil, err
	}
	serverServer, err := server.Provide(serverDiConfig, brokerBroker, loggerLogger, registryRegistry, transportTransport, serverConfig)
	if err != nil {
		return nil, err
	}
	service, err := provideService(opts, brokerBroker, clientClient, loggerLogger, registryRegistry, serverServer, transportTransport)
	if err != nil {
		return nil, err
	}
	return service, nil
}
