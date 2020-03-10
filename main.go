package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "dex-cli",
		Usage: "Manage Dex OIDC provider.",
		Commands: []*cli.Command{
			{
				Name: "client",
			},
			{
				Name: "password",
			},
			{
				Name: "refresh",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
