package fs

func IsFolderEmpty(folder Folder) bool {
	return len(folder.List()) == 0
}

func GetFolderList(folder Folder) []string {
	var folderList []string
	fileList := folder.List()
	for _, info := range fileList {
		if !info.IsFile() {
			folderList = append(folderList, info.Name())
		}
	}
	return folderList
}

func GetFileList(folder Folder) []string {
	var folderList []string
	fileList := folder.List()
	for _, info := range fileList {
		if info.IsFile() {
			folderList = append(folderList, info.Name())
		}
	}
	return folderList
}
