package fs

import "github.com/AldieNightStar/mhistea_go/_common"

type folderInfoImp struct {
	name   string
	isFile bool
}

func NewFolderInfo(name string, isFile bool) _common.FolderInfo {
	return &folderInfoImp{
		name:   name,
		isFile: isFile,
	}
}

func (f *folderInfoImp) Name() string {
	return f.name
}

func (f *folderInfoImp) IsFile() bool {
	return f.isFile
}
