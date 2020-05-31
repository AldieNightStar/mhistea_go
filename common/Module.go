package common

type Module interface {
	Name() string
	Script() string
	Config() string
}
