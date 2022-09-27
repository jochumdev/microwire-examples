package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-micro/microwire"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	app, err := microwire.DefaultApp(
		mWire.Name("livecyclehooks"),
		mWire.Usage("A POC for go-micro.dev/v5"),
		mWire.Version("v0.0.1"),
		mWire.ArgPrefix(""),
		mWire.Action(func(cli *cli.Context, service micro.Service) error {
			fmt.Println("Action executed")
			return nil
		}),
		mWire.BeforeStart(func() error {
			fmt.Println("BeforeStart")
			return nil
		}),
		mWire.BeforeStop(func() error {
			fmt.Println("BeforeStop")
			return nil
		}),
		mWire.AfterStart(func() error {
			fmt.Println("AfterStart")
			return nil
		}),
		mWire.AfterStop(func() error {
			fmt.Println("AfterStop")
			return errors.New("failure on stop?")
		}),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
