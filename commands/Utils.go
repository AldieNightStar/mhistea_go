package commands

func AddCommandWithSynonyms(registry CommandRegistry, synonyms []string, moduleName, alias string) {
	if len(synonyms) == 0 {
		return
	}
	for _, name := range synonyms {
		registry.AddCommand(name, moduleName, alias)
	}
}
