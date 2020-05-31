package sections

import "github.com/AldieNightStar/mhistea_go/_common"

type sectionReaderObject struct{}

func NewSectionReader() _common.SectionReader {
	return sectionReaderObject{}
}

func (s sectionReaderObject) GetSectionList(text string) []string {
	var list []string
	sections := GetSections(text)
	for _, section := range sections {
		list = append(list, section.Name)
	}
	return list
}

func (s sectionReaderObject) GetSectionByName(text, name string) string {
	return ReadSectionByName(text, name)
}
