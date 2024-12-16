package handlers

import (
	"context"
	"fmt"

	"github.com/nifle3/gosay/internal/message"
	"github.com/urfave/cli/v3"
)

func Say(_ context.Context, cmd *cli.Command) error {
	sayMessage := message.GenerateSay(10, cmd.Args().Slice())
	fmt.Print(sayMessage)

	return nil
}
