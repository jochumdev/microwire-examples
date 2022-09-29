package main

import (
	"errors"
	"fmt"

	"github.com/go-micro/microwire"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	_ "github.com/go-micro/microwire/plugins/cli/urfave"
	_ "github.com/go-micro/microwire/plugins/transport/http"
	_ "github.com/go-micro/plugins/v4/broker/http"
	_ "github.com/go-micro/plugins/v4/registry/mdns"
)

func main() {
	service, err := microwire.NewService(
		microwire.Name("livecyclehooks"),
		microwire.Usage("A POC for go-micro.dev/v5"),
		microwire.Version("v0.0.1"),
		microwire.ArgPrefix(""),
		microwire.Action(func(config microwire.Store, service micro.Service) error {
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

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
