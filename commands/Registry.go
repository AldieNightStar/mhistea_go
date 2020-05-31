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

//	Create new command registry
func NewCommandRegistry() common.CommandRegistry {
	return &Commands{
		m: make(map[string]CommandInfo),
	}
}

//	Adds command to registry
//		commandName - Name of the command. Examples: "print" or "pr" etc
//		moduleName  - Module name in which this command exists
//		alias		- Real module function which will be called when command get used
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

func (c *Commands) CommandList() []string {
	var list []string
	for k, _ := range c.m {
		list = append(list, k)
	}
	return list
}

func modCall(moduleName, funcName, args string) string {
	return "mods[\"" + moduleName + "\"][\"" + funcName + "\"](" + args + ");"
}
