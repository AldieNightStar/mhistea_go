package fs

import (
	"io/ioutil"
	"os"
	"strings"
)

type LocalFs struct {
	Path string
}

func (fs LocalFs) List() []FolderInfo {
	files, err := ioutil.ReadDir(fs.Path)
	if err != nil {
		return nil
	}
	var list []FolderInfo
	for i := 0; i < len(files); i++ {
		file := files[i]
		info := FolderInfo{
			file.Name(),
			!file.IsDir(),
		}
		list = append(list, info)
	}
	return list
}

func (fs LocalFs) ReadFile(name string) []byte {
	path := suffixSlash(fs.Path)
	data, err := ioutil.ReadFile(path + name)
	if err != nil {
		return []byte{}
	}
	return data
}

func (fs LocalFs) WriteFile(name string, data []byte) bool {
	path := suffixSlash(fs.Path)
	err := ioutil.WriteFile(path+name, data, 0)
	if err != nil {
		return false
	}
	return true
}

func (fs LocalFs) GetFolder(name string) Folder {
	nextPath := suffixSlash(fs.Path) + name
	_, err := ioutil.ReadDir(nextPath)
	if err != nil {
		return nil
	}
	newFolder := &LocalFs{
		Path: nextPath,
	}
	return newFolder
}

func (fs LocalFs) IsFile(name string) bool {
	info, err := os.Stat(suffixSlash(fs.Path) + name)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func (fs LocalFs) DeleteFile(name string) bool {
	filePath := suffixSlash(fs.Path) + name
	err := os.Remove(filePath)
	if err != nil {
		return false
	}
	return true
}

func (fs LocalFs) DeleteFolder(name string) bool {
	return fs.DeleteFile(name)
}

func (fs LocalFs) CreateFolder(name string) bool {
	err := os.Mkdir(name, 0)
	if err != nil {
		return false
	}
	return true
}

func (fs LocalFs) IsExists(name string) bool {
	_, err := os.Stat(suffixSlash(fs.Path) + name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func NewLocalFolder(path string) Folder {
	return &LocalFs{Path: path}
}

func suffixSlash(path string) string {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}
