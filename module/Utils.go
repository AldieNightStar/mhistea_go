package module

import "github.com/AldieNightStar/mhistea_go/_common"

// Read module configuration file as [SectionalConfiguration] object
func ReadConfig(mod _common.Module) _common.SectionalConfiguration {
	cfgText := mod.Config()
	parseConfig := _common.Refs.Config.ParseConfig
	if parseConfig == nil {
		return nil
	}
	return parseConfig(cfgText)
}

func ReadWrappedScript(mod _common.Module) string {
	parseTemplate := _common.Refs.Parser.ParseTemplate
	if parseTemplate == nil {
		return ""
	}
	script := mod.Script()
	return parseTemplate(wrappedModScript, []string{mod.Name(), script})
}

func ReadCommands(mod _common.Module) []_common.CommandInfo {
	parseConfig := _common.Refs.Config.ParseConfig
	if parseConfig == nil {
		return nil
	}
	cfgText := mod.Config()
	cfg := parseConfig(cfgText)

	commandKeys := cfg.Keys("commands")
	if commandKeys == nil {
		return nil
	}
	var list []_common.CommandInfo
	for _, cmdKey := range commandKeys {
		alias := cfg.Get("commands", cmdKey)
		list = append(list, _common.CommandInfo{
			Name:   cmdKey,
			Alias:  alias,
			Module: mod.Name(),
		})
	}
	return list
}
