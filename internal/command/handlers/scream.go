package handlers

import (
	"context"
	"fmt"

	"github.com/nifle3/gosay/internal/message"
	"github.com/urfave/cli/v3"
)

func Scream(_ context.Context, cmd *cli.Command) error {
	sayMessage := message.GenerateScream(10, cmd.Args().Slice())
	fmt.Print(sayMessage)

	return nil
}
