package main

import (
	"github.com/go-micro/microwire/microwire"
	"github.com/go-micro/microwire/util/cmd"
	"github.com/urfave/cli/v2"
)

func ProvideBrokerFlags(opts *microwire.Options) microwire.BrokerFlags {
	return microwire.BrokerFlags{
		&cli.StringFlag{
			Name:    cmd.PrefixName(opts.ArgPrefix, "broker"),
			Usage:   "Broker for pub/sub. http, nats, rabbitmq",
			Value:   "nats",
			EnvVars: []string{cmd.PrefixEnv(opts.ArgPrefix, "BROKER")},
		},
		&cli.StringFlag{
			Name:    cmd.PrefixName(opts.ArgPrefix, "broker_address"),
			Usage:   "Comma-separated list of broker addresses",
			EnvVars: []string{cmd.PrefixEnv(opts.ArgPrefix, "BROKER_ADDRESS")},
		},
	}
}
