package _common

type SectionalConfiguration interface {
	Get(section, key string) (value string)
	Sections() []string
	Keys(section string) []string
}
