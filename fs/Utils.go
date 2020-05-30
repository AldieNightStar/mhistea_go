package fs

import "github.com/AldieNightStar/mhistea_go/common"

func IsFolderEmpty(folder common.Folder) bool {
	return len(folder.List()) == 0
}

func GetFolderList(folder common.Folder) []string {
	var folderList []string
	fileList := folder.List()
	for _, info := range fileList {
		if !info.IsFile() {
			folderList = append(folderList, info.Name())
		}
	}
	return folderList
}

func GetFileList(folder common.Folder) []string {
	var folderList []string
	fileList := folder.List()
	for _, info := range fileList {
		if info.IsFile() {
			folderList = append(folderList, info.Name())
		}
	}
	return folderList
}

// Recursive delete file or directory.
// Even if directory is not empty
func RecursiveDelete(fs common.Folder, name string) bool {
	if !fs.IsExists(name) {
		return false
	}
	if fs.IsFile(name) {
		return fs.DeleteFile(name)
	}
	folder := fs.GetFolder(name)
	if folder == nil {
		return false
	}
	list := folder.List()
	for _, info := range list {
		status := RecursiveDelete(folder, info.Name())
		if !status {
			return false
		}
	}
	fs.DeleteFolder(name)
	return true
}
