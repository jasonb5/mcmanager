package main

import (
	"fmt"
	"log"
	"mcmanager/internal/curse"
	"mcmanager/internal/mcmanager"
)

func main() {
	config := curse.DefaultSearchConfig
	config.SearchFilter = "Skyopolis"

	data, _ := curse.Search(config, nil)

	for i, x := range data {
		fmt.Printf("%d\t%d\t%v\n", i, x.ID, x.Name)
	}

	versions, _ := data[1].GetVersions()

	for i, x := range versions {
		fmt.Printf("%d\t%d\t%v\n", i, x.ID, x.DisplayName)
	}

	installConfig := mcmanager.NewConfig("sky4", "test")

	// if err := curse.InstallServer(&data[1], &versions[14], installConfig, nil); err != nil {
	// 	log.Fatalf("error %v", err)
	// }

	if err := curse.InstallServerByID(433578, 3474592, installConfig, nil); err != nil {
		log.Fatalf("error %v", err)
	}
}
