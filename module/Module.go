package module

import (
	"errors"
	"github.com/AldieNightStar/mhistea_go/common"
)

type modImpl struct {
	folder common.ReadOnlyFolder
	name   string
}

func (m modImpl) Name() string {
	return m.name
}

func (m modImpl) Script() string {
	return string(m.folder.ReadFile("mod.js"))
}

func (m modImpl) Config() string {
	return string(m.folder.ReadFile("config.txt"))
}

func NewModule(name string, folder common.Folder) (mod common.Module, err error) {
	if !folder.IsExists("mod.js") {
		return nil, errors.New("mod.js is not present!")
	}
	if !folder.IsExists("config.txt") {
		return nil, errors.New("config.txt is not present!")
	}
	return &modImpl{folder, name}, nil
}
