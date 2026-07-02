package cli

type Command interface {
	Execute(args []string) (string, error)
}
