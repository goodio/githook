// +build darwin dragonfly freebsd linux netbsd openbsd

package main

import (
	"flag"
	"os"
	"fmt"
	"os/exec"
)

const DEFAULT_STORE = `/tmp`
var godaemon = flag.Bool("d", true, "run app as a daemon with -d=true")

func pull(gitRoot string) error {
	if gitRoot == "" {
		return fmt.Errorf("%s", "gitRoot is empty!!")
	}

	if _, err := runCommand(gitRoot, "/bin/sh", "-c", "git fetch --all && git reset --hard origin/master && git pull"); err != nil {
		return err
	}

	return nil
}

func init() {

	flag.Parse()

	if *godaemon {
		args := os.Args[1:]
		i := 0
		for ; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()

		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}