package cmd

import (
	"easymod/internal/ui"

	"github.com/urfave/cli/v2"
)

// Execute serves as the cli entry point
func Execute() *cli.App {
	app := &cli.App{
		Name:     "easymod",
		Usage:    "Easily build file permissions without remembering chmod syntax.",
		Commands: []*cli.Command{},
		Action: func(c *cli.Context) error {

			if err := ui.InitScreen(); err != nil {
				return err
			}

			return nil
		},
	}

	return app
}
