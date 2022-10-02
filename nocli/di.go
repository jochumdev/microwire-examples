package main

import (
	"github.com/go-micro/microwire"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/transport"
)

func provideService(
	opts *microwire.Options,
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

	service := micro.NewService(
		mOpts...,
	)

	for _, fn := range opts.Actions {
		if err := fn(service); err != nil {
			return nil, err
		}
	}

	return service, nil
}
