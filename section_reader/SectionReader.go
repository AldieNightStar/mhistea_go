package section_reader

import "strings"

type SectionInfo struct {
	Name       string
	LineNumber int32
}

const NOT_FOUND = -1
const SECTION_PREFIX = "::"

func ReadSectionByName(text, sectionName string) (out string) {
	if strings.Contains(text, "\r") {
		text = strings.ReplaceAll(text, "\r", "")
	}
	sections := GetSections(text)
	if len(sections) == 0 {
		return ""
	}
	lines := SplitToLines(text)
	lineNumber := int32(NOT_FOUND)
	for i := 0; i < len(sections); i++ {
		section := sections[i]
		if section.Name == sectionName {
			lineNumber = section.LineNumber + 1
			break
		}
	}
	if lineNumber == NOT_FOUND {
		return ""
	}
	sb := strings.Builder{}
	for i := int(lineNumber); i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, SECTION_PREFIX) {
			break
		}
		sb.WriteString(line + "\n")
	}
	return sb.String()
}

func GetSections(text string) (sections []SectionInfo) {
	sections = []SectionInfo{}
	if text == "" {
		return sections
	}
	if strings.Contains(text, "\r") {
		text = strings.ReplaceAll(text, "\r", "")
	}
	if !strings.Contains(text, "\n") {
		return sections
	}
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, SECTION_PREFIX) {
			prefixLen := len(SECTION_PREFIX)
			sectionName := strings.Trim(line[prefixLen:], " ")
			sections = append(sections, SectionInfo{
				Name:       sectionName,
				LineNumber: int32(i),
			})
		}
	}
	return sections
}

func SplitToLines(text string) (lines []string) {
	if text == "" {
		return []string{}
	}
	if !strings.Contains(text, "\n") {
		return []string{text}
	}
	return strings.Split(text, "\n")
}
