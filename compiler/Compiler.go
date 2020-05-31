package compiler

import (
	"errors"
	"github.com/AldieNightStar/mhistea_go/_common"
	"strings"
)

//	Compile story into HTML file
//		template - Template which will be used to put story script to
//		modulesFolder - Folder with modules (Folders which has "config.txt" and "mod.js")
//		storyFolder - Folder which has "config.txt" and "story.txt" files. Also "res" folder with resources
//			In this folder will be writed compile story "index.html"
func CompileStory(templatesFolder _common.ReadOnlyFolder, modulesFolder, storyFolder _common.Folder) error {
	// Import functions from [common] package
	NewCommandRegistry := _common.Refs.Commands.NewCommandRegistry
	NewStory := _common.Refs.Story.NewStory
	ParseConfig := _common.Refs.Config.ParseConfig
	CommaSplit := _common.Refs.Commands.CommaSplit
	NewModule := _common.Refs.Module.NewModule
	ReadCommands := _common.Refs.Module.ReadCommands
	ReadWrappedScript := _common.Refs.Module.ReadWrappedScript
	Pack := _common.Refs.Packer.Pack

	// Check args for nil's
	if templatesFolder == nil {
		return errors.New("Template folder is nil!")
	}
	if modulesFolder == nil {
		return errors.New("Modules Folder is nil!")
	}
	if storyFolder == nil {
		return errors.New("story Folder is nil!")
	}

	// Create Story object
	story, err := NewStory(storyFolder)
	if err != nil {
		return err
	}

	// Create command registry
	cmdRegistry := NewCommandRegistry()

	// Read story config
	storyConfig := ParseConfig(story.StoryConfig())

	// Read modules to connect
	storyModuleNames := CommaSplit(storyConfig.Get("config", "mods"))
	{
		if len(storyModuleNames) == 0 {
			return errors.New("At least one module has to be in module list!")
		}
	}

	// Load modules
	var modules []_common.Module
	{
		for _, modName := range storyModuleNames {
			mod, err := NewModule(modulesFolder.GetFolder(modName), modName)
			if err != nil || mod == nil {
				return errors.New("Module [" + modName + "] Failed to load or missing!")
			}
			modules = append(modules, mod)
		}
	}

	// Aggregate module commands
	{
		for _, mod := range modules {
			commands := ReadCommands(mod)
			if len(commands) == 0 {
				continue
			}
			for _, cmd := range commands {
				cmdRegistry.AddCommand(cmd.Name, cmd.Module, cmd.Alias)
			}
		}
	}

	// Concatenate modules into one definition script
	var modulesOutputScript string
	{
		builder := strings.Builder{}
		for _, mod := range modules {
			wscript := ReadWrappedScript(mod)
			builder.WriteString(wscript)
		}
		modulesOutputScript = modulesDefinition + scenesDefinition + builder.String()
	}

	// Load story text
	storyText := story.StoryText()

	// Compile story to script
	storyOutputScript := compileStoryText(storyText, cmdRegistry)

	// Concatenate Story output and Modules output scripts into one script
	outputScript := modulesOutputScript + storyOutputScript

	// Load template
	var templateText string
	{
		templateName := storyConfig.Get("config", "template")
		if templateName == "" {
			return errors.New("Config parameter 'template' is empty!")
		}
		if !templatesFolder.IsExists(templateName + ".html") {
			return errors.New("Template is not exists: " + templateName + ".html")
		}
		templateText = string(templatesFolder.ReadFile(templateName + ".html"))
	}

	// Pack story output script into template
	var packedHTML string
	{
		html, err := Pack(templateText, outputScript)
		if err != nil {
			return err
		}
		packedHTML = html
	}

	// Write PackedStory into file "index.html" inside folder
	{
		const outputStoryHtmlFileName = "index.html"
		exist := storyFolder.IsExists(outputStoryHtmlFileName)
		if exist {
			ok := storyFolder.DeleteFile(outputStoryHtmlFileName)
			if !ok {
				return errors.New("Can't delete file " + outputStoryHtmlFileName + " to create new one!")
			}
		}
		storyFolder.WriteFile(outputStoryHtmlFileName, []byte(packedHTML))
	}

	// Return nil instead of error, which say that compilation is successful
	return nil
}

func compileStoryText(storyText string, cmdRegistry _common.CommandRegistry) string {
	// Imports
	SectionReader := _common.Refs.Sections.SectionReader
	ParseTemplate := _common.Refs.Parser.ParseTemplate

	builder := strings.Builder{}

	sections := SectionReader.GetSectionList(storyText)
	for _, section := range sections {
		parsedSection := compileSectionText(SectionReader.GetSectionByName(storyText, section), cmdRegistry)
		parsedSection = ParseTemplate(sectionFunctionWrapper, []string{section, parsedSection})
		builder.WriteString(parsedSection)
	}

	return builder.String()
}

func compileSectionText(storyText string, cmdRegistry _common.CommandRegistry) string {
	SplitToLines := _common.Refs.Sections.SplitToLines
	ParseCommandAndArguments := _common.Refs.Parser.ParseCommandAndArguments

	storyTextLines := SplitToLines(storyText)
	compiledStoryScript := strings.Builder{}
	for _, storyTextLine := range storyTextLines {
		// Skip empty lines and comments ("==")
		if storyTextLine == "" || strings.HasPrefix(storyTextLine, "==") {
			continue
		}
		if strings.HasPrefix(storyTextLine, "\t") || strings.HasPrefix(storyTextLine, "  ") {
			storyTextLine = strings.Trim(storyTextLine, " \t")
			compiledStoryScript.WriteString(storyTextLine)
			compiledStoryScript.WriteString("\n")
			continue
		} else if strings.HasPrefix(storyTextLine, ".") {
			storyTextLine = storyTextLine[1:]
			cmd, arg := ParseCommandAndArguments(storyTextLine)
			usedCommand := cmdRegistry.UseCommand(cmd, arg)
			compiledStoryScript.WriteString(usedCommand)
			compiledStoryScript.WriteString("\n")
		} else {
			if strings.Contains(storyTextLine, "`") {
				storyTextLine = strings.ReplaceAll(storyTextLine, "`", "\"")
			}
			if strings.Contains(storyTextLine, "\\") {
				storyTextLine = strings.ReplaceAll(storyTextLine, "\\", "\\\\")
			}
			if strings.Contains(storyTextLine, "$") {
				storyTextLine = strings.ReplaceAll(storyTextLine, "$", "\\$")
			}
			usedCommand := cmdRegistry.UseCommand("print", "`"+storyTextLine+"`")
			if usedCommand != "" {
				compiledStoryScript.WriteString(usedCommand)
				compiledStoryScript.WriteString("\n")
			}
		}
	}
	return compiledStoryScript.String()
}
