package main

import (
	"github.com/reimersoftware/dex-cli/dex"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cli.HelpFlag = &cli.BoolFlag{
		Name: "help", Aliases: []string{"h", "?"},
		Usage: "Show help",
	}
	app := &cli.App{
		Name:  "dex-cli",
		Usage: "Manage Dex OIDC provider.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "target",
				Aliases:  []string{"t"},
				Usage:    "Dex server `ADDRESS`",
				EnvVars:  []string{"TARGET", "DEX_TARGET"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "client",
				Aliases: []string{"c"},
				Usage:   "Manage client applications",
				Subcommands: []*cli.Command{
					{
						Name:  "create",
						Usage: "Create a new client application",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "id",
								Usage:    "Unique `ID`",
								Aliases:  []string{"i"},
								Required: true,
							},
							&cli.StringSliceFlag{
								Name:     "redirect-uris",
								Usage:    "`URIS…` to redirect to after successful authorization",
								Aliases:  []string{"r"},
								Required: true,
							},
							&cli.StringSliceFlag{
								Name:    "trusted-peers",
								Usage:   "Trusted `PEERS…`",
								Aliases: []string{"t"},
							},
							&cli.BoolFlag{
								Name:    "public",
								Usage:   "The client secret should not be kept private",
								Aliases: []string{"p"},
								Value:   false,
							},
							&cli.StringFlag{
								Name:     "name",
								Usage:    "Human-readable `NAME`",
								Aliases:  []string{"n"},
								Required: true,
							},
							&cli.StringFlag{
								Usage:   "Logo `URL`",
								Name:    "logo-url",
								Aliases: []string{"l"},
							},
						},
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.CreateClient(
									context.String("id"),
									context.StringSlice("redirect-uris"),
									context.StringSlice("trusted-peers"),
									context.Bool("public"),
									context.String("name"),
									context.String("logo-url"),
								)
							})
						},
					},
					{
						Name:  "delete",
						Usage: "Delete an existing client application",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "id",
								Usage:    "Unique `IDENTIFIER`",
								Aliases:  []string{"i"},
								Required: true,
							},
						},
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.DeleteClient(
									context.String("id"),
								)
							})
						},
					},
				},
			},
			{
				Name:    "user",
				Aliases: []string{"p"},
				Usage:   "Manage user passwords",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "List all users.",
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.ListPasswords()
							})
						},
					},
					{
						Name:  "create",
						Usage: "Create a new user",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "email",
								Usage:    "`EMAIL` address",
								Aliases:  []string{"e"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "username",
								Usage:    "Public user `NAME`",
								Aliases:  []string{"u"},
								Required: true,
							},
							&cli.StringFlag{
								Name:    "password",
								Usage:   "It is recommended to set the `PASSWORD` from a file or the standard input stream",
								Aliases: []string{"p"},
								EnvVars: []string{"PASSWORD", "DEX_PASSWORD"},
							},
							&cli.StringFlag{
								Name:    "password-file",
								Usage:   "Read password from `FILE`",
								Aliases: []string{"f"},
							},
							&cli.BoolFlag{
								Name:    "password-stdin",
								Usage:   "Read password from standard input stream",
								Aliases: []string{"s"},
							},
						},
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.CreatePassword(
									context.String("email"),
									context.String("username"),
									context.String("password"),
									context.String("password-text"),
									context.Bool("password-stdin"),
								)
							})
						},
					},
					{
						Name:  "update",
						Usage: "Update an existing user",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "email",
								Usage:    "`EMAIL` address (cannot be changed)",
								Aliases:  []string{"e"},
								Required: true,
							},
							&cli.StringFlag{
								Name:    "username",
								Usage:   "Public user `NAME`",
								Aliases: []string{"u"},
							},
							&cli.StringFlag{
								Name:    "password",
								Usage:   "It is recommended to set the `PASSWORD` from a file or the standard input stream",
								Aliases: []string{"p"},
								EnvVars: []string{"PASSWORD", "DEX_PASSWORD"},
							},
							&cli.PathFlag{
								Name:    "password-file",
								Usage:   "Read password from `FILE`",
								Aliases: []string{"f"},
							},
							&cli.BoolFlag{
								Name:    "password-stdin",
								Usage:   "Read password from standard input stream",
								Aliases: []string{"s"},
							},
						},
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.UpdatePassword(
									context.String("email"),
									context.String("username"),
									context.String("password"),
									context.String("password-text"),
									context.Bool("password-stdin"),
								)
							})
						},
					},
					{
						Name:  "delete",
						Usage: "Delete an existing user",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "email",
								Usage:    "`EMAIL` address",
								Aliases:  []string{"e"},
								Required: true,
							},
						},
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.DeletePassword(
									context.String("email"),
								)
							})
						},
					},
				},
			},
			{
				Name:    "refresh",
				Aliases: []string{"r"},
				Usage:   "Manage refresh tokens",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Usage: "List user's refresh tokens",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "user",
								Usage:    "User `UUID`",
								Aliases:  []string{"u"},
								Required: true,
							},
						},
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.ListRefresh(
									context.String("user"),
								)
							})
						},
					},
					{
						Name:  "revoke",
						Usage: "Revoke a user's refresh token for a client application",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "user",
								Usage:    "User `UUID`",
								Aliases:  []string{"u"},
								Required: true,
							},
							&cli.StringFlag{
								Name:     "client",
								Usage:    "Client `ID`",
								Aliases:  []string{"c"},
								Required: true,
							},
						},
						Action: func(context *cli.Context) error {
							return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
								return dex.RevokeRefresh(
									context.String("user"),
									context.String("client"),
								)
							})
						},
					},
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show Dex server version",
				Action: func(context *cli.Context) error {
					return dex.UseDex(context.String("target"), func(dex dex.Dex) error {
						return dex.GetVersion()
					})
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.New(os.Stderr, "", 0).Fatal(err)
	}
}
