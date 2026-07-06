package commands

import (
	"strings"
)

type EchoCommand struct{}

func (cmd EchoCommand) Execute(args []string) (string, error) {
	return strings.Join(args, " "), nil
}
