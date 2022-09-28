//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/go-micro/microwire"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"go-micro.dev/v4"
)

func NewWireService(opts ...mWire.Option) (micro.Service, error) {
	panic(wire.Build(
		microwire.ProvideOptions,
		microwire.ProvideCLI,
		microwire.ProvideCliArgs,

		ProvideBrokerFlags,
		ProvideRegistryFlags,
		ProvideTransportFlags,
		ProvideBrokerOpts,
		ProvideRegistryOpts,
		ProvideTransportOpts,

		microwire.AllComponentsSet,
		microwire.ProvideInitializedCLI,
		microwire.ProvideMicroOpts,
		microwire.ProvideMicroService,
	))
}
