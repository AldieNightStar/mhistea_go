package config

type SectionReader interface {
	GetSectionList(text string) []string
	GetSectionByName(text, name string) string
}

type Configuration interface {
	Get(section, key string) (value string)
}

type FileReader interface {
	ReadFile(fileName string) []byte
}
