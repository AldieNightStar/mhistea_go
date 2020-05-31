package _init

import (
	"github.com/AldieNightStar/mhistea_go/_common"
	"github.com/AldieNightStar/mhistea_go/commands"
	"github.com/AldieNightStar/mhistea_go/config"
	"github.com/AldieNightStar/mhistea_go/parser"
	"github.com/AldieNightStar/mhistea_go/sections"
)

func Init() {
	// Commands
	_common.Refs.Commands.AddCommandWithSynonymCommaSeparated = commands.AddCommandWithSynonymCommaSeparated
	_common.Refs.Commands.CommaSplit = commands.CommaSplit
	// Config
	_common.Refs.Config.ParseConfig = config.ParseConfig
	// Sections
	_common.Refs.Sections.SectionReader = sections.NewSectionReader()
	// Parser
	_common.Refs.Parser.ParseTemplate = parser.ParseTemplate
}
