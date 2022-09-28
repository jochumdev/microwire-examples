package main

import (
	"errors"
	"fmt"

	"github.com/go-micro/microwire"
	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/microwire/plugins/cli/urfave"
	_ "github.com/go-micro/microwire/plugins/transport/http"
	_ "github.com/go-micro/plugins/v4/broker/http"
	_ "github.com/go-micro/plugins/v4/registry/mdns"
)

func main() {
	service, err := microwire.WireService(
		mWire.Name("livecyclehooks"),
		mWire.Usage("A POC for go-micro.dev/v5"),
		mWire.Version("v0.0.1"),
		mWire.ArgPrefix(""),
		mWire.Action(func(cli mCli.CLI, service micro.Service) error {
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

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
