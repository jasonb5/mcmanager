package download

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WriteEntry(f *zip.File, path string) error {
	outputPath := filepath.Join(path, f.Name)

	if strings.HasSuffix(f.Name, "/") {
		if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory %v: %v", outputPath, err)
		}
	} else {
		if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory %v: %v", filepath.Dir(outputPath), err)
		}

		outputFile, err := os.Create(outputPath)

		if err != nil {
			return fmt.Errorf("error creating output file %v: %v", outputPath, err)
		}

		defer outputFile.Close()

		inputFile, err := f.Open()

		if err != nil {
			return fmt.Errorf("error reading entry %v: %v", f.Name, err)
		}

		defer inputFile.Close()

		if _, err := io.Copy(outputFile, inputFile); err != nil {
			return fmt.Errorf("error coping zip entry to file %v: %v", outputPath, err)
		}
	}

	return nil
}

func Unzip(zipPath, extractPath string) error {
	reader, err := zip.OpenReader(zipPath)

	if err != nil {
		return fmt.Errorf("error opening zip %v: %v", zipPath, err)
	}

	defer reader.Close()

	for _, f := range reader.File {
		if err := WriteEntry(f, extractPath); err != nil {
			return fmt.Errorf("error extracting %v: %v", f.Name, err)
		}
	}

	log.Printf("Extracted %d entries", len(reader.File))

	return nil
}
