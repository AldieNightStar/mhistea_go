package _common

type ConfigModule interface {
	Config() string
}

type ScriptingModule interface {
	Script() string
}

type Module interface {
	Name() string
	ScriptingModule
	ConfigModule
}
