package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mcmanager/internal/curse"
	"mcmanager/internal/mcmanager"
	"os"
	"sort"
	"strconv"

	"github.com/urfave/cli/v2"
)

func GetenvInt(name string) int {
	value, err := strconv.Atoi(os.Getenv(name))

	if err != nil {
		log.Fatalf("error could not convert %v to int", os.Getenv(name))
	}

	return value
}

func searchModPacks(c *cli.Context) error {
	config := curse.DefaultSearchConfig
	config.SearchFilter = c.String("filter")

	packs, err := curse.Search(config, nil)

	if err != nil {
		return err
	}

	fmt.Print("ID\tName\n")

	for _, p := range packs {
		fmt.Printf("%d\t%s\n", p.ID, p.Name)
	}

	return nil
}

type byName []curse.ModPackVersion

func (s byName) Len() int {
	return len(s)
}

func (s byName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byName) Less(i, j int) bool {
	return s[i].DisplayName < s[j].DisplayName
}

func searchVersions(c *cli.Context) error {
	modpack, err := curse.GetModPack(c.Int("id"), nil)

	if err != nil {
		return err
	}

	versions, err := modpack.GetVersions()

	if err != nil {
		return err
	}

	sort.Sort(byName(versions))

	for _, v := range versions {
		fmt.Printf("%d\t%s\n", v.ID, v.DisplayName)
	}

	return nil
}

func runModPack(c *cli.Context) error {
	config := mcmanager.NewConfig(c.String("name"), c.String("path"))

	if err := curse.InstallServerByID(c.Int("id"), c.Int("versionID"), config, nil); err != nil {
		return err
	}

	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "search",
				Usage: "search for modpacks",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "filter",
						Usage: "filter search query",
					},
				},
				Action: searchModPacks,
			},
			{
				Name:  "versions",
				Usage: "list versions for modpack",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "id",
						Usage:    "modpack id to retrieve versions for",
						Required: true,
					},
				},
				Action: searchVersions,
			},
			{
				Name:  "run",
				Usage: "install and run a modpack",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Usage:   "name of the sub-directory that the modpack will be installed in",
						EnvVars: []string{"MCMANAGER_NAME"},
					},
					&cli.StringFlag{
						Name:    "path",
						Usage:   "path to where the modpack will be installed",
						EnvVars: []string{"MCMANAGER_PATH"},
						Value:   ".",
					},
					&cli.IntFlag{
						Name:     "id",
						Usage:    "modpack id to install",
						Required: true,
						EnvVars:  []string{"MCMANAGER_MODPACK_ID"},
					},
					&cli.IntFlag{
						Name:     "versionID",
						Usage:    "version of modpack to install",
						Required: true,
						EnvVars:  []string{"MCMANAGER_VERSION_ID"},
					},
				},
				Action: runModPack,
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "log"},
		},
		Before: func(c *cli.Context) error {
			if !c.Bool("log") {
				log.SetOutput(ioutil.Discard)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
