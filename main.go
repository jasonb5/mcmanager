package main

import (
	"fmt"
	"log"
	"mcmanager/internal/curse"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "installPath",
				Usage:   "Path to install Modpacks under",
				Value:   ".",
				EnvVars: []string{"MCM_INSTALL_PATH"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "curse",
				Usage: "Install a curse modpack",
				Subcommands: []*cli.Command{
					{
						Name:    "search",
						Aliases: []string{"s"},
						Usage:   "Search for a modpack",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "filter",
								Usage: "Filter to search with e.g. skyopolis",
							},
						},
						Action: func(c *cli.Context) error {
							config := curse.DefaultConfig
							config.SearchFilter = c.String("filter")

							packs, err := curse.Search(config)

							if err != nil {
								return err
							}

							for _, item := range packs {
								fmt.Println(item.ID, item.Name)
							}

							return nil
						},
					},
					{
						Name:    "version",
						Aliases: []string{"v"},
						Usage:   "List version for a modpack",
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:     "id",
								Usage:    "Modpack ID to find versions for",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							modPackID := c.Int("id")

							versions, err := curse.GetVersions(modPackID)

							if err != nil {
								return err
							}

							for _, item := range versions {
								if item.ServerPackFileID == 0 {
									continue
								}

								fmt.Println(item.ServerPackFileID, item.DisplayName)
							}

							return nil
						},
					},
					{
						Name:    "run",
						Aliases: []string{"r"},
						Usage:   "Install and run modpack",
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:     "id",
								Usage:    "Modpack ID",
								Required: true,
								EnvVars:  []string{"MCM_ID"},
							},
							&cli.IntFlag{
								Name:     "version",
								Usage:    "Version ID",
								Required: true,
								EnvVars:  []string{"MCM_VERSION"},
							},
						},
						Action: func(c *cli.Context) error {
							installPath := c.String("installPath")
							modPackID := c.Int("id")
							versionID := c.Int("version")

							installPath, err := filepath.Abs(installPath)

							if err != nil {
								return err
							}

							dataPath, err := curse.Install(modPackID, versionID, installPath)

							if err != nil {
								return err
							}

							if err := curse.Run(dataPath); err != nil {
								return err
							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
