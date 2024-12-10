package commands

type CommandInterface interface {
	Name() string
	Execute() (string, error)
	CheckArgs(args CommandArguments) bool
	SetArgs(args CommandArguments)
}
