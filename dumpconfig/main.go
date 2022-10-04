package main

import (
	"fmt"
	"os"

	_ "github.com/go-micro/microwire-plugins/client/grpc/v5"
	_ "github.com/go-micro/microwire-plugins/logger/logrus/v5"
	_ "github.com/go-micro/microwire-plugins/logger/zap/v5"
	_ "github.com/go-micro/microwire-plugins/server/grpc/v5"
	_ "github.com/go-micro/microwire-plugins/sets/defaults/v5"
	_ "github.com/go-micro/microwire-plugins/sets/nats/v5"
	_ "github.com/go-micro/microwire-plugins/transport/quic/v5"
	micro "github.com/go-micro/microwire/v5"
	"github.com/go-micro/microwire/v5/logger"
)

func main() {
	service, err := micro.NewService(
		micro.Name("dumpconfig"),
		micro.Usage("A POC for go-micro.dev/v5"),
		micro.Version("v0.0.1"),
		micro.ArgPrefix(""),
		micro.BeforeStart(func(service micro.Service) error {
			b, err := service.DumpConfig()
			if err != nil {
				return err
			}

			fmt.Printf("\n\n---\n%s\n", string(b))

			os.Exit(0)
			return nil
		}),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
