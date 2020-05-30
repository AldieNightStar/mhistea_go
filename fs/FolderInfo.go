package fs

import "github.com/AldieNightStar/mhistea_go/common"

type FolderInfoImp struct {
	name   string
	isFile bool
}

func NewFolderInfo(name string, isFile bool) common.FolderInfo {
	return &FolderInfoImp{
		name:   name,
		isFile: isFile,
	}
}

func (f *FolderInfoImp) Name() string {
	return f.name
}

func (f *FolderInfoImp) IsFile() bool {
	return f.isFile
}
