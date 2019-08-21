// +build windows

package githook

import (
	"fmt"
	"github.com/ghaoo/githook/utils"
)

const DEFAULT_STORE = "C:/HookTemp"

func Pull(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	// 开发环境，谨慎拉取
	if _, err := utils.RunCommand(gitRoot, "git", "pull"); err != nil {
		return err
	}

	return nil
}

func Push(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	if _, err := utils.RunCommand(gitRoot, "/bin/sh", "-c", "git push -u -f origin master"); err != nil {
		return err
	}

	return nil
}
