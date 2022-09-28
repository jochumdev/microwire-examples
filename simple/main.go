package main

import (
	"github.com/go-micro/microwire"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/microwire/plugins/cli/urfave"
	_ "github.com/go-micro/microwire/plugins/transport/http"
	_ "github.com/go-micro/plugins/v4/broker/http"
	_ "github.com/go-micro/plugins/v4/registry/mdns"
)

func main() {
	service, err := microwire.NewWireService(
		mWire.Name("livecyclehooks"),
		mWire.Usage("A POC for go-micro.dev/v5"),
		mWire.Version("v0.0.1"),
		mWire.ArgPrefix(""),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
