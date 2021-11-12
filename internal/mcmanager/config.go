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
	installPath := filepath.Join(path, name)

	installPath, err := filepath.Abs(installPath)

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
