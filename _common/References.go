package _common

type RefStruct struct {
	Commands struct {
		//	from [commands.AddCommandWithSynonymCommaSeparated]
		AddCommandWithSynonymCommaSeparated func(registry CommandRegistry, synonyms, moduleName, alias string)
		//	from [commands.CommaSplit]
		CommaSplit func(str string) []string
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
	}
	Parser struct {
		//	from [parser.ParseTemplate]
		ParseTemplate func(template string, text []string) (out string)
	}
}

var Refs = &RefStruct{}
