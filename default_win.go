// +build windows

package main

import "fmt"

const DEFAULT_STORE = "C:/HookTemp"

func pull(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	// 开发环境，谨慎拉取
	if _, err := runCommand(gitRoot, "git", "pull"); err != nil {
		return err
	}

	return nil
}

func push(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	if _, err := runCommand(gitRoot, "/bin/sh", "-c", "git push -u origin master -f"); err != nil {
		return err
	}

	return nil
}
