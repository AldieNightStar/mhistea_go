package common

type SectionalConfiguration interface {
	Get(section, key string) (value string)
}
