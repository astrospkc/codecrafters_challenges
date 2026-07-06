package commands

import "os"

type ExitCommand struct{}

func (cmd ExitCommand) Execute(args []string) (string, error) {
	os.Exit(0)
	return "", nil
}
