package commands

import "github.com/AldieNightStar/mhistea_go/common"

func AddCommandWithSynonyms(registry common.CommandRegistry, synonyms []string, moduleName, alias string) {
	if len(synonyms) == 0 {
		return
	}
	for _, name := range synonyms {
		registry.AddCommand(name, moduleName, alias)
	}
}
