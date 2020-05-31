package config

import (
	"github.com/AldieNightStar/mhistea_go/common"
	"strings"
)

//	Reading configuration from [FileReader] and parse by [SectionReader]
//	Returns Configuration
//		cfg := ReadConfig(fileReader, sectionReader, "config.txt")
//		cfg.Get("section_a", "name") // etc
func ReadConfig(fileReader common.FileReader, sectionReader common.SectionReader, fileName string) common.SectionalConfiguration {
	data := fileReader.ReadFile(fileName)
	if data == nil || len(data) == 0 {
		return nil
	}
	text := string(data)
	return ParseConfig(sectionReader, text)
}

//	Parse configuration text and returns [SectionalConfiguration]
//	Example:
//		:: Profile
//		name = Antonio
//		age = 25
func ParseConfig(sectionReader common.SectionReader, text string) common.SectionalConfiguration {
	sectionList := sectionReader.GetSectionList(text)
	if len(sectionList) == 0 {
		return nil
	}
	config := cfg{
		m: make(map[string]map[string]string),
	}
	for _, sectionName := range sectionList {
		sectionText := sectionReader.GetSectionByName(text, sectionName)
		sectionConfig := parseConfigSection(sectionText)
		config.m[sectionName] = sectionConfig
	}
	return config
}

func parseConfigSection(text string) map[string]string {
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
		if strings.Contains(key, ",") {
			synonyms := parseSynonyms(key)
			for _, synonym := range synonyms {
				m[synonym] = val
			}
		} else {
			m[key] = val
		}
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

func parseSynonyms(key string) []string {
	if key == "" {
		return nil
	}
	if !strings.Contains(key, ",") {
		return []string{key}
	}
	array := strings.Split(key, ",")
	var list []string
	for _, subKey := range array {
		subKey = strings.Trim(subKey, " \t")
		list = append(list, subKey)
	}
	return list
}
