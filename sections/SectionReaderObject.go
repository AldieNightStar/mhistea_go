package sections

type SectionReaderObject struct{}

func (s SectionReaderObject) GetSectionList(text string) []string {
	var list []string
	sections := GetSections(text)
	for _, section := range sections {
		list = append(list, section.Name)
	}
	return list
}

func (s SectionReaderObject) GetSectionByName(text, name string) string {
	return ReadSectionByName(text, name)
}
