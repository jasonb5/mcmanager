package ftb

import (
	"bufio"
	"log"
	"os/exec"
	"path/filepath"
)

func Run(path string) error {
	startPath := filepath.Join(path, "start.sh")

	cmd := exec.Command(startPath)

	cmd.Dir = path

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
