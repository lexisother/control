package lib

import (
	"bufio"
	"fmt"
	"os/exec"
)

func RunCmd(dir string, command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	err := cmd.Start()

	errScanner := bufio.NewScanner(stderr)
	errScanner.Split(bufio.ScanLines)
	for errScanner.Scan() {
		m := errScanner.Text()
		fmt.Println(m)
	}

	outScanner := bufio.NewScanner(stdout)
	outScanner.Split(bufio.ScanLines)
	for outScanner.Scan() {
		m := outScanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()

	if err != nil {
		return err
	}
	return nil
}
