package _init

import (
	"github.com/AldieNightStar/mhistea_go/_common"
	"github.com/AldieNightStar/mhistea_go/commands"
	"github.com/AldieNightStar/mhistea_go/config"
	"github.com/AldieNightStar/mhistea_go/module"
	"github.com/AldieNightStar/mhistea_go/packer"
	"github.com/AldieNightStar/mhistea_go/parser"
	"github.com/AldieNightStar/mhistea_go/sections"
	"github.com/AldieNightStar/mhistea_go/story"
)

func Init() {
	// Commands
	_common.Refs.Commands.AddCommandWithSynonymCommaSeparated = commands.AddCommandWithSynonymCommaSeparated
	_common.Refs.Commands.CommaSplit = commands.CommaSplit
	_common.Refs.Commands.NewCommandRegistry = commands.NewCommandRegistry
	// Config
	_common.Refs.Config.ParseConfig = config.ParseConfig
	// Sections
	_common.Refs.Sections.SectionReader = sections.NewSectionReader()
	_common.Refs.Sections.SplitToLines = sections.SplitToLines
	// Parser
	_common.Refs.Parser.ParseTemplate = parser.ParseTemplate
	_common.Refs.Parser.ParseCommandAndArguments = parser.ParseCommandAndArguments
	// Packer
	_common.Refs.Packer.Pack = packer.Pack
	// Story
	_common.Refs.Story.NewStory = story.NewStory
	// Modules
	_common.Refs.Module.NewModule = module.NewModule
	_common.Refs.Module.ReadCommands = module.ReadCommands
	_common.Refs.Module.ReadWrappedScript = module.ReadWrappedScript
}
