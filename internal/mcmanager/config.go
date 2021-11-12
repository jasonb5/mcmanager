package mcmanager

import (
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Name, InstallPath string
}

func NewConfig(name, path string) *Config {
	installPath, err := filepath.Abs(path)

	if err != nil {
		log.Fatalf("error getting absolute path for %v: %v", installPath, err)
	}

	if err := os.MkdirAll(installPath, os.ModePerm); err != nil {
		log.Fatalf("error creating install directory %v: %v", installPath, err)
	}

	return &Config{
		Name:        name,
		InstallPath: installPath,
	}
}
