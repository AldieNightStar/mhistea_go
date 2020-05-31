package module

import "github.com/AldieNightStar/mhistea_go/common"

func ReadConfig(mod common.Module, cfgParser common.ConfigParser, sReader common.SectionReader) common.SectionalConfiguration {
	cfgText := mod.Config()
	return cfgParser(sReader, cfgText)
}

func ReadWrappedScript(mod common.Module, parse common.ParseTemplate) string {
	script := mod.Script()
	return parse(wrappedModScript, []string{mod.Name(), script})
}
