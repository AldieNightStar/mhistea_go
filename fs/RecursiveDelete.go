package fs

// Recursive delete file or directory.
// Even if directory is not empty
func RecursiveDelete(fs Folder, name string) bool {
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
		status := RecursiveDelete(folder, info.Name)
		if !status {
			return false
		}
	}
	fs.DeleteFolder(name)
	return true
}
