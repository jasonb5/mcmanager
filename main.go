package main

import (
	"log"
	"mcmanager/internal/curse"
	"mcmanager/internal/mcmanager"
	"os"
	"strconv"
)

func GetenvInt(name string) int {
	value, err := strconv.Atoi(os.Getenv(name))

	if err != nil {
		log.Fatalf("error could not convert %v to int", os.Getenv(name))
	}

	return value
}

func main() {
	installPath := os.Getenv("MCMANAGER_PATH")

	name := os.Getenv("MCMANAGER_NAME")

	modpackID := GetenvInt("MCMANAGER_MODPACK_ID")

	versionID := GetenvInt("MCMANAGER_VERSION_ID")

	config := mcmanager.NewConfig(name, installPath)

	if err := curse.InstallServerByID(modpackID, versionID, config, nil); err != nil {
		log.Fatalf("could not install curse server pack: %v", err)
	}
}
