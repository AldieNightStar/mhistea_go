package common

type FileSection interface {
	GetSectionList(text string) []string
	GetSectionByName(text, name string) string
}
