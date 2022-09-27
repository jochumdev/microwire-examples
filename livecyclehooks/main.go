package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-micro/microwire/microwire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	app, err := microwire.DefaultApp(
		microwire.Name("livecyclehooks"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.ArgPrefix(""),
		microwire.Action(func(cli *cli.Context, service micro.Service) error {
			fmt.Println("Action executed")
			return nil
		}),
		microwire.BeforeStart(func() error {
			fmt.Println("BeforeStart")
			return nil
		}),
		microwire.BeforeStop(func() error {
			fmt.Println("BeforeStop")
			return nil
		}),
		microwire.AfterStart(func() error {
			fmt.Println("AfterStart")
			return nil
		}),
		microwire.AfterStop(func() error {
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
