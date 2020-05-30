package config

import "strings"

//	Configuration
//		cfg := ReadConfig(fileReader, sectionReader, "config.txt")
//		cfg.Get("section_a", "name") // etc
type Config struct {
	m map[string]map[string]string
}

func (c Config) Get(section, key string) (value string) {
	return c.m[section][key]
}

//	Reading configuration from [FileReader] and parse by [SectionReader]
//	Returns Configuration
//		cfg := ReadConfig(fileReader, sectionReader, "config.txt")
//		cfg.Get("section_a", "name") // etc
func ReadConfig(fileReader FileReader, sectionReader SectionReader, fileName string) Configuration {
	data := fileReader.ReadFile(fileName)
	if data == nil || len(data) == 0 {
		return nil
	}
	text := string(data)
	sectionList := sectionReader.GetSectionList(text)
	if len(sectionList) == 0 {
		return nil
	}
	config := Config{
		m: make(map[string]map[string]string),
	}
	for _, sectionName := range sectionList {
		sectionText := sectionReader.GetSectionByName(text, sectionName)
		sectionConfig := ParseConfig(sectionText)
		config.m[sectionName] = sectionConfig
	}
	return config
}

//	Parse configuration text
//	Example:
//		name = Antonio
//		age = 25
//	Will not parse ::sections
//	Good to use for each ::section to have MultiConfig
func ParseConfig(text string) map[string]string {
	if text == "" {
		return nil
	}
	m := make(map[string]string)
	if !strings.Contains(text, "\n") {
		key, val := parseLine(text)
		m[key] = val
		return m
	}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		key, val := parseLine(line)
		m[key] = val
	}
	return m
}

func parseLine(text string) (key, value string) {
	if !strings.Contains(text, "=") {
		return text, "true"
	}
	arr := strings.SplitN(text, "=", 2)
	arr[0] = strings.Trim(arr[0], " \t")
	arr[1] = strings.Trim(arr[1], " \t")
	return arr[0], arr[1]
}
