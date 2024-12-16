package main

import (
	"context"
	"log"
	"os"

	"github.com/nifle3/gosay/internal/command"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			command.Say,
			command.Scream,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
