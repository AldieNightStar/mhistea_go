package commands

import (
	"errors"
	"github.com/AldieNightStar/mhistea_go/common"
)

type CommandInfo struct {
	Alias, Module string
}

type Commands struct {
	m map[string]CommandInfo
}

func NewCommandRegistry() common.CommandRegistry {
	return &Commands{
		m: make(map[string]CommandInfo),
	}
}

func (c *Commands) AddCommand(commandName, moduleName, alias string) error {
	_, exist := c.m[commandName]
	if exist {
		return errors.New("Such command already exists: " + commandName)
	}
	c.m[commandName] = CommandInfo{
		Alias:  alias,
		Module: moduleName,
	}
	return nil
}

func (c *Commands) UseCommand(commandName, args string) string {
	info, ok := c.m[commandName]
	if !ok {
		return ""
	}
	return modCall(info.Module, info.Alias, args)
}

func (c *Commands) GetCommandAlias(commandName string) string {
	return c.m[commandName].Alias
}

func (c *Commands) GetCommandModuleName(commandName string) string {
	return c.m[commandName].Module
}

func modCall(moduleName, funcName, args string) string {
	return "mods[\"" + moduleName + "\"][\"" + funcName + "\"](" + args + ");"
}
