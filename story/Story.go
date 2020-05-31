package story

import (
	"errors"
	"github.com/AldieNightStar/mhistea_go/common"
)

type story struct {
	folder common.ReadOnlyFolder
}

func (s story) ResourceList() []string {
	res := s.folder.GetFolder("res")
	if res == nil {
		return nil
	}
	resList := res.List()
	var list []string
	for _, res := range resList {
		if res.IsFile() {
			list = append(list, res.Name())
		}
	}
	return list
}

func (s story) StoryText() string {
	return string(s.folder.ReadFile("story.txt"))
}

func (s story) Resource(name string) []byte {
	res := s.folder.GetFolder("res")
	if res == nil {
		return nil
	}
	return res.ReadFile(name)
}

func (s story) StoryConfig() string {
	return string(s.folder.ReadFile("config.txt"))
}

func NewStory(folder common.ReadOnlyFolder) (common.Story, error) {
	if !folder.IsExists("story.txt") {
		return nil, errors.New("Story have no story.txt file!")
	}
	if !folder.IsExists("config.txt") {
		return nil, errors.New("Story have no config.txt file!")
	}
	if !folder.IsExists("res") || !folder.IsExists("res") {
		return nil, errors.New("Story have no [res] folder!")
	}
	return &story{folder: folder}, nil
}
