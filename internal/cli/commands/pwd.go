package commands

import (
	"fmt"
	"os"
)

type PwdCommand struct{}

func (cmd PwdCommand) Execute(args []string) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf("pwd: too many arguments")
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("pwd: %v", err)
	}

	return dir, nil
}
