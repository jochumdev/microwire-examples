package main

import (
	"os"

	"github.com/go-micro/microwire"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4/logger"
)

func main() {
	app, err := microwire.DefaultApp(
		mWire.Name("livecyclehooks"),
		mWire.Usage("A POC for go-micro.dev/v5"),
		mWire.Version("v0.0.1"),
		mWire.ArgPrefix(""),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
