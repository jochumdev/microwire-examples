package main

import (
	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/transport"
)

func provideService(
	opts *micro.Options,
	broker broker.Broker,
	registry registry.Registry,
	transport transport.Transport,
) (micro.Service, error) {
	mOpts := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
	}

	if broker != nil {
		mOpts = append(mOpts, micro.Broker(broker))
	}
	if registry != nil {
		mOpts = append(mOpts, micro.Registry(registry))
	}
	if transport != nil {
		mOpts = append(mOpts, micro.Transport(transport))
	}

	service := micro.NewMicroService(
		mOpts...,
	)

	return service, nil
}
