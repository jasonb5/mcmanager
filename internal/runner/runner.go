package runner

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(execPath string) error {
	cmd := exec.Command(execPath)

	log.Printf("Running command %v", execPath)

	basePath := filepath.Dir(execPath)

	if err := os.Chdir(basePath); err != nil {
		return fmt.Errorf("error changing directory to %v: %v", basePath, err)
	}

	stderr, err := cmd.StderrPipe()

	if err != nil {
		return fmt.Errorf("error getting strerr for process: %v", err)
	}

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return fmt.Errorf("error getting stdout for process: %v", err)
	}

	multi := io.MultiReader(stdout, stderr)

	scanner := bufio.NewScanner(multi)

	go func() {
		for scanner.Scan() {
			msg := scanner.Text()

			log.Print(msg)
		}
	}()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running process: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting on process: %v", err)
	}

	return nil
}
