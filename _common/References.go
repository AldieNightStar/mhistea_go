package _common

type RefStruct struct {
	Commands struct {
		//	from [commands.AddCommandWithSynonymCommaSeparated]
		AddCommandWithSynonymCommaSeparated func(registry CommandRegistry, synonyms, moduleName, alias string)
		//	from [commands.CommaSplit]
		CommaSplit func(str string) []string
		//	from [commands.NewCommandRegistry]
		NewCommandRegistry func() CommandRegistry
	}
	Config struct {
		//	from [config.ParseConfig]
		ParseConfig func(text string) SectionalConfiguration
		//	from [config.ReadConfig]
		ReadConfig func(fileReader FileReader, fileName string) SectionalConfiguration
	}
	Sections struct {
		//	from [sections.SectionReader]
		SectionReader SectionReader
		//	from [sections.SplitToLines]
		SplitToLines func(text string) (lines []string)
	}
	Parser struct {
		//	from [parser.ParseTemplate]
		ParseTemplate func(template string, text []string) (out string)
		//	from [parser.ParseCommandAndArguments]
		ParseCommandAndArguments func(line string) (command, args string)
	}
	Packer struct {
		//	from [packer.Pack]
		Pack func(template, scriptBundle string) (out string, err error)
	}
	Story struct {
		//	from [story.NewStory]
		NewStory func(folder ReadOnlyFolder) (Story, error)
	}
	Module struct {
		// from [module.NewModule]
		NewModule func(folder ReadOnlyFolder, name string) (mod Module, err error)

		// from [module.ReadCommands]
		ReadCommands func(mod Module) []CommandInfo

		// from [module.ReadWrappedScript]
		ReadWrappedScript func(mod Module) string
	}
	Subtext struct {
		// from [subtext.ReadSubtext]
		ReadSubtext func(text string) []string
	}
}

var Refs = &RefStruct{}
