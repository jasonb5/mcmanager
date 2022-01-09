package ftb

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func convertSlug(name string) string {
	safeName := strings.ReplaceAll(name, " ", "_")

	return strings.ToLower(safeName)
}

func getVersion(term string) (*Version, *Manifest, error) {
	modPacks, err := Search(term)

	if err != nil {
		return nil, nil, err
	}

	if len(modPacks) > 1 {
		return nil, nil, fmt.Errorf("found %d modpacks, refine search", len(modPacks))
	}

	var version *Version

	modPack := modPacks[0]

	// find latest version
	for _, item := range modPack.Versions {
		if item.Type != "Release" {
			continue
		}

		if version == nil {
			version = &item
		} else {
			if item.ID > version.ID {
				version = &item
			}
		}
	}

	return version, modPack, nil
}

func dumpVersionMetadata(version *Version, path string) error {
	versionPath := filepath.Join(path, "version.json")

	data, err := json.MarshalIndent(version, "", "  ")

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(versionPath, data, fs.ModePerm); err != nil {
		return err
	}

	return nil
}

func downloadServer(modPack, version int, path string) (string, error) {
	downloadURL := fmt.Sprintf("https://api.modpacks.ch/public/modpack/%d/%d/server/linux", modPack, version)

	res, err := http.Get(downloadURL)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("serverinstall_%d_%d", modPack, version)

	installerPath := filepath.Join(path, filename)

	if err := ioutil.WriteFile(installerPath, data, fs.ModePerm); err != nil {
		return "", err
	}

	if err := os.Chmod(installerPath, 0755); err != nil {
		return "", err
	}

	return installerPath, nil
}

func runInstaller(installerPath, dataPath string) error {
	installPath := filepath.Dir(installerPath)

	cmd := exec.Command(installerPath, "--auto", "--path", dataPath)

	cmd.Dir = installPath

	messages := make(chan string)

	stdOut, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	stdOutScanner := bufio.NewScanner(stdOut)

	go func() {
		for stdOutScanner.Scan() {
			text := stdOutScanner.Text()

			messages <- text
		}
	}()

	stdErr, err := cmd.StderrPipe()

	if err != nil {
		return err
	}

	stdErrScanner := bufio.NewScanner(stdErr)

	go func() {
		for stdErrScanner.Scan() {
			text := stdErrScanner.Text()

			messages <- text
		}
	}()

	go func() {
		for text := range messages {
			log.Print(text)
		}
	}()

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func writeEULA(path string) error {
	eulaPath := filepath.Join(path, "eula.txt")

	if err := ioutil.WriteFile(eulaPath, []byte("eula=true"), fs.ModePerm); err != nil {
		return err
	}

	return nil
}

func Install(term, path string) (string, error) {
	version, modPack, err := getVersion(term)

	if err != nil {
		return "", err
	}

	slug := convertSlug(modPack.Name)

	installPath := filepath.Join(path, slug)

	if err := os.MkdirAll(installPath, fs.ModePerm); err != nil {
		return "", err
	}

	if err := dumpVersionMetadata(version, installPath); err != nil {
		return "", err
	}

	installerPath, err := downloadServer(modPack.ID, version.ID, installPath)

	if err != nil {
		return "", err
	}

	dataPath := filepath.Join(installPath, "data")

	if err := os.MkdirAll(dataPath, fs.ModePerm); err != nil {
		return "", err
	}

	if err := runInstaller(installerPath, dataPath); err != nil {
		return "", err
	}

	if err := writeEULA(dataPath); err != nil {
		return "", err
	}

	return dataPath, nil
}
