// +build darwin dragonfly freebsd linux netbsd openbsd

package main

import (
	"fmt"
)

const DEFAULT_STORE = `/tmp`

func pull(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	// 线上环境，强制覆盖
	if _, err := runCommand(gitRoot, "/bin/sh", "-c", "git fetch --all && git reset --hard origin/master && git pull"); err != nil {
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