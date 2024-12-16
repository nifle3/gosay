package command

import (
	"github.com/nifle3/gosay/internal/command/handlers"
	"github.com/urfave/cli/v3"
)

var DefaultCommand = "say"

var Say = &cli.Command{
	Name:   DefaultCommand,
	Usage:  "Gopher is talk to you",
	Action: handlers.Say,
}

var Scream = &cli.Command{
	Name:   "scream",
	Usage:  "Gopher is scream to you",
	Action: handlers.Scream,
}
