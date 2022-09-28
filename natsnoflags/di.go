package main

import (
	mBroker "github.com/go-micro/microwire/broker"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
)

// Stubbed out
func ProvideBrokerFlags() *mBroker.DiFlags       { return nil }
func ProvideRegistryFlags() *mRegistry.DiFlags   { return nil }
func ProvideTransportFlags() *mTransport.DiFlags { return nil }

func ProvideBrokerOpts(opts *mWire.Options, c mWire.InitializedCli) (*mBroker.DiOptions, error) {
	return &mBroker.DiOptions{
		Plugin:    "nats",
		Addresses: "nats://localhost:4222",
	}, nil
}

func ProvideRegistryOpts(opts *mWire.Options, c mWire.InitializedCli) (*mRegistry.DiOptions, error) {
	return &mRegistry.DiOptions{
		Plugin:    "nats",
		Addresses: "nats://localhost:4222",
	}, nil
}

func ProvideTransportOpts(opts *mWire.Options, c mWire.InitializedCli) (*mTransport.DiOptions, error) {
	return &mTransport.DiOptions{
		Plugin:    "nats",
		Addresses: "nats://localhost:4222",
	}, nil
}
