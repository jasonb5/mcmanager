package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func writeFile(file *zip.File, outputPath string) error {
	outputPath = filepath.Join(outputPath, file.Name)

	if strings.HasSuffix(file.Name, "/") {
		if err := os.MkdirAll(outputPath, file.Mode()); err != nil {
			return err
		}
	} else {
		if err := os.MkdirAll(filepath.Dir(outputPath), file.Mode()); err != nil {
			return err
		}

		outputFile, err := os.Create(outputPath)

		if err != nil {
			return err
		}

		defer outputFile.Close()

		inputFile, err := file.Open()

		if err != nil {
			return err
		}

		defer inputFile.Close()

		if _, err := io.Copy(outputFile, inputFile); err != nil {
			return err
		}

		log.Printf("Wrote %s -> %s", file.Name, outputPath)

		if err := os.Chmod(outputPath, file.Mode()); err != nil {
			return err
		}
	}

	return nil
}

func Unzip(src, dst string) error {
	log.Printf("Unzipping %s -> %s", src, dst)

	archive, err := zip.OpenReader(src)

	if err != nil {
		return err
	}

	defer archive.Close()

	for _, f := range archive.File {
		if err := writeFile(f, dst); err != nil {
			return nil
		}
	}

	log.Printf("Done unzipping %d entries", len(archive.File))

	return nil
}
