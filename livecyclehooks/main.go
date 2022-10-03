package main

import (
	"errors"
	"fmt"

	micro "github.com/go-micro/microwire/v5"

	"github.com/go-micro/microwire/v5/logger"

	_ "github.com/go-micro/microwire-plugins/sets/defaults/v5"
)

func main() {
	service, err := micro.NewService(
		micro.Name("livecyclehooks"),
		micro.Usage("A POC for go-micro.dev/v5"),
		micro.Version("v0.0.1"),
		micro.ArgPrefix(""),
		micro.Action(func(service micro.Service) error {
			fmt.Println("Action executed")
			return nil
		}),
		micro.BeforeStart(func(_ micro.Service) error {
			fmt.Println("BeforeStart")
			return nil
		}),
		micro.BeforeStop(func(_ micro.Service) error {
			fmt.Println("BeforeStop")
			return nil
		}),
		micro.AfterStart(func(_ micro.Service) error {
			fmt.Println("AfterStart")
			return nil
		}),
		micro.AfterStop(func(_ micro.Service) error {
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
