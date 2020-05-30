package common

type SectionReader interface {
	GetSectionList(text string) []string
	GetSectionByName(text, name string) string
}
