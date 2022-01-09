package curse

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

func Run(dataPath string) error {
	serverStarterConfig := path.Join(dataPath, "server-setup-config.yaml")

	if _, err := os.Stat(serverStarterConfig); err == nil {
		log.Printf("Found ServerStarter config")

		serverStarterScript := path.Join(dataPath, "startserver.sh")

		if err := os.Chmod(serverStarterScript, 0700); err != nil {
			return err
		}

		cmd := exec.Command(serverStarterScript)

		cmd.Dir = dataPath

		cmd.Env = os.Environ()

		stdErr, err := cmd.StderrPipe()

		if err != nil {
			return err
		}

		stdOut, err := cmd.StdoutPipe()

		if err != nil {
			return err
		}

		multiReader := io.MultiReader(stdOut, stdErr)

		scanner := bufio.NewScanner(multiReader)

		go func() {
			log.Println("Scanning stdout and stderr")

			for scanner.Scan() {
				text := scanner.Text()

				log.Print(text)
			}

			log.Println("Done scanning stdout and stderr")
		}()

		log.Println("Starting process")

		if err := cmd.Run(); err != nil {
			return err
		}

		log.Println("Waiting on process")

		if err := cmd.Wait(); err != nil {
			return err
		}

		log.Println("Done waiting on process")
	} else {
		return fmt.Errorf("standard server run not implemented")
	}

	return nil
}
