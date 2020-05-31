package fs

import (
	"github.com/AldieNightStar/mhistea_go/_common"
	"io/ioutil"
	"os"
	"strings"
)

type localFs struct {
	Path string
}

func (fs localFs) List() []_common.FolderInfo {
	files, err := ioutil.ReadDir(fs.Path)
	if err != nil {
		return nil
	}
	var list []_common.FolderInfo
	for i := 0; i < len(files); i++ {
		file := files[i]
		info := NewFolderInfo(file.Name(), !file.IsDir())
		list = append(list, info)
	}
	return list
}

func (fs localFs) ReadFile(name string) []byte {
	path := suffixSlash(fs.Path)
	data, err := ioutil.ReadFile(path + name)
	if err != nil {
		return []byte{}
	}
	return data
}

func (fs localFs) WriteFile(name string, data []byte) bool {
	path := suffixSlash(fs.Path)
	err := ioutil.WriteFile(path+name, data, 0)
	if err != nil {
		return false
	}
	return true
}

func (fs localFs) GetFolder(name string) _common.Folder {
	nextPath := suffixSlash(fs.Path) + name
	_, err := ioutil.ReadDir(nextPath)
	if err != nil {
		return nil
	}
	newFolder := &localFs{
		Path: nextPath,
	}
	return newFolder
}

func (fs localFs) IsFile(name string) bool {
	info, err := os.Stat(suffixSlash(fs.Path) + name)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func (fs localFs) DeleteFile(name string) bool {
	filePath := suffixSlash(fs.Path) + name
	err := os.Remove(filePath)
	if err != nil {
		return false
	}
	return true
}

func (fs localFs) DeleteFolder(name string) bool {
	return fs.DeleteFile(name)
}

func (fs localFs) CreateFolder(name string) bool {
	err := os.Mkdir(name, 0)
	if err != nil {
		return false
	}
	return true
}

func (fs localFs) IsExists(name string) bool {
	_, err := os.Stat(suffixSlash(fs.Path) + name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func NewLocalFolder(path string) _common.Folder {
	return &localFs{Path: path}
}

func suffixSlash(path string) string {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}
