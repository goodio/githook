// +build windows

package main

import "fmt"

const DEFAULT_STORE = "C:/HookTemp"

func pull(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	if _, err := runCommand(gitRoot, "git", "pull"); err != nil {
		return err
	}

	return nil
}
