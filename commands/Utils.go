package commands

import (
	"github.com/AldieNightStar/mhistea_go/_common"
	"strings"
)

func AddCommandWithSynonyms(registry _common.CommandRegistry, synonyms []string, moduleName, alias string) {
	if len(synonyms) == 0 {
		return
	}
	for _, name := range synonyms {
		registry.AddCommand(name, moduleName, alias)
	}
}

func AddCommandWithSynonymCommaSeparated(registry _common.CommandRegistry, synonyms, moduleName, alias string) {
	list := CommaSplit(synonyms)
	AddCommandWithSynonyms(registry, list, moduleName, alias)
}

func CommaSplit(str string) []string {
	if str == "" {
		return []string{}
	}
	if !strings.Contains(str, ",") {
		return []string{str}
	}
	arr := strings.Split(str, ",")
	var list []string
	for _, elem := range arr {
		list = append(list, strings.Trim(elem, " \t"))
	}
	return list
}
