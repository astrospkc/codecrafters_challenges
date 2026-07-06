package commands

import (
	"fmt"
	"os"
)

const HomeDir string = "~"

type CdCommand struct{}

func (cmd CdCommand) Execute(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("type requires 1 argument")
	}

	path := args[0]
	if path == HomeDir {
		path = os.Getenv("HOME")
	}

	err := os.Chdir(path)
	if err != nil {
		return "", fmt.Errorf("cd: %s: No such file or directory", args[0])
	}
	return "", nil
}
