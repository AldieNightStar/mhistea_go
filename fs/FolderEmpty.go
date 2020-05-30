package fs

func IsFolderEmpty(folder Folder) bool {
	return len(folder.List()) == 0
}
