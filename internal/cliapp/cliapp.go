package cliapp

import (
	"code"
	"context"
	"fmt"
	"io"

	"github.com/urfave/cli/v3"
)

const pathArgumentError = "exactly two path arguments are required"

func New(stdout io.Writer) *cli.Command {
	return &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		ArgsUsage: "<filepath1> <filepath2>",
		Writer:    stdout,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			args := cmd.Args()
			if args.Len() != 2 {
				return cli.Exit(pathArgumentError, 1)
			}

			filepath1 := args.Get(0)
			filepath2 := args.Get(1)
			diff, err := code.GenDiff(filepath1, filepath2, cmd.String("format"))
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			_, err = fmt.Fprintf(cmd.Root().Writer, "%s\n", diff)
			return err
		},
	}
}
