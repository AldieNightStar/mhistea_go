package _common

type ResourceableStory interface {
	ResourceList() []string
	Resource(name string) []byte
}

type SimpleStory interface {
	StoryText() string
}

type ConfigurableStory interface {
	StoryConfig() string
}

type Story interface {
	ResourceableStory
	SimpleStory
	ConfigurableStory
}
