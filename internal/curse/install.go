package curse

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"mcmanager/internal/utils"
	"os"
	"path"
	"path/filepath"
)

func dumpVersionMetadata(version *Version, path string) error {
	data, err := json.MarshalIndent(version, "", "  ")

	if err != nil {
		return err
	}

	metadataFilePath := filepath.Join(path, "version.json")

	if err := ioutil.WriteFile(metadataFilePath, data, fs.ModePerm); err != nil {
		return err
	}

	return nil
}

func writeEULA(path string) error {
	eulaFile := filepath.Join(path, "eula.txt")

	if err := ioutil.WriteFile(eulaFile, []byte("eula=true"), fs.ModePerm); err != nil {
		return err
	}

	return nil
}

func Install(modPackID, versionID int, installPath string) (string, error) {
	modPack, err := GetModpack(modPackID)

	if err != nil {
		return "", err
	}

	version, err := GetVersionByID(modPackID, versionID)

	if err != nil {
		return "", err
	}

	modPackPath := filepath.Join(installPath, modPack.Slug)

	log.Printf("Installing %s to %s", modPack.Name, modPackPath)

	if err := os.MkdirAll(modPackPath, os.ModePerm); err != nil {
		return "", err
	}

	if err := dumpVersionMetadata(version, modPackPath); err != nil {
		return "", err
	}

	modPackArchive, err := DownloadMod(modPack.ID, versionID, modPackPath)

	if err != nil {
		return "", err
	}

	dataPath := path.Join(modPackPath, "data")

	if err := os.MkdirAll(dataPath, os.ModePerm); err != nil {
		return "", err
	}

	if err := utils.Unzip(modPackArchive, dataPath); err != nil {
		return "", err
	}

	if err := writeEULA(dataPath); err != nil {
		return "", err
	}

	return dataPath, nil
}
