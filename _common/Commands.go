package _common

type CommandRegistry interface {
	AddCommand(commandName, moduleName, alias string) error
	UseCommand(commandName, args string) string
	GetCommandAlias(commandName string) string
	GetCommandModuleName(commandName string) string
	CommandList() []string
}

type CommandInfo struct {
	Name, Alias, Module string
}
