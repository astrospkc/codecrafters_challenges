package commands

import (
	"errors"
	"fmt"
	"os/exec"
)

type TypeCommand struct {
	IsBuiltin func(name string) bool
}

func (cmd TypeCommand) Execute(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("type requires 1 argument")
	}

	name := args[0]
	if cmd.IsBuiltin(name) {
		return fmt.Sprintf("%s is a shell built-in", name), nil
	}
	if path, err := exec.LookPath(name); err == nil {
		return fmt.Sprintf("%s is %s", name, path), nil
	}
	return "", fmt.Errorf("%s: not found", name)
}
