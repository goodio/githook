package main

import (
	"fmt"
	"os/exec"
)

func runCommand(cwd, command string, args ...string) (string, error) {

	cmd := exec.Command(command, args...)
	if cwd != "" {
		cmd.Dir = cwd
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command: %v: %q", err, string(output))
	}

	return string(output), nil
}
