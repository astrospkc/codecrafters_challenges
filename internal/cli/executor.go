package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Executor struct {
	registry *Registry
}

func NewExecutor(registry *Registry) *Executor {
	return &Executor{
		registry: registry,
	}
}

func openOrCreateFile(redirection *Redirection) (*os.File, error) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	return os.OpenFile(redirection.target, flags, 0644)
}

func (e *Executor) Execute(command PCommand) error {
	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr

	if command.Redirection != nil {
		f, err := openOrCreateFile(command.Redirection)
		if err != nil {
			return err
		}
		defer f.Close()

		switch command.Redirection.fd {
		case -1, 1:
			stdout = f
			break
		case 2:
			stderr = f
		}
	}

	if len(command.Args) == 0 {
		return nil
	}

	commandExecutable, exists := e.registry.Commands[command.Args[0]]

	if exists {
		if output, err := commandExecutable.Execute(command.Args[1:]); err != nil {
			return err
		} else if output != "" {
			output += "\n"
			_, err := stdout.Write([]byte(output))
			if err != nil {
				return err
			}
		}
	} else {
		pathCommand := exec.Command(command.Args[0], command.Args[1:]...)
		pathCommand.Stdout = stdout
		pathCommand.Stderr = stderr
		pathCommand.Stdin = os.Stdin

		if err := pathCommand.Run(); err != nil {
			if errors.Is(err, exec.ErrNotFound) {
				return fmt.Errorf("%s: not found", command.Args[0])
			}
			//return err
		}
	}

	return nil
}
