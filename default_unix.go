// +build darwin dragonfly freebsd linux netbsd openbsd

package githook

import (
	"fmt"
	"github.com/ghaoo/githook/utils"
)

const DEFAULT_STORE = `/tmp`

func Pull(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	// 线上环境，强制覆盖
	if _, err := utils.RunCommand(gitRoot, "/bin/sh", "-c", "git fetch --all && git reset --hard origin/master && git pull"); err != nil {
		return err
	}

	return nil
}

func Push(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	if _, err := utils.RunCommand(gitRoot, "/bin/sh", "-c", "git push -u origin master -f"); err != nil {
		return err
	}

	return nil
}