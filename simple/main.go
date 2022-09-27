package main

import (
	"os"

	"github.com/go-micro/microwire/microwire"
	"go-micro.dev/v4/logger"
)

func main() {
	app, err := microwire.DefaultApp(
		microwire.Name("livecyclehooks"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.ArgPrefix(""),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
