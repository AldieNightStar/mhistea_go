package fs

type Folder interface {
	// List of files and folders
	List() []FolderInfo
	// Read file
	ReadFile(name string) []byte
	// Write File (Rewrite if exists)
	WriteFile(name string, data []byte) bool
	// Get Folder
	// Returns same structure of Folder or nil if there is no such
	GetFolder(name string) Folder
	// Delete File
	DeleteFile(name string) bool
	// Delete Folder
	// true - folder deleted successfully
	// false - Folder is not removable or it is not empty
	DeleteFolder(name string) bool
	// Create Folder
	CreateFolder(name string) bool
	// Check isFile
	IsFile(name string) bool
	// Check if file exist
	IsExists(name string) bool
}

type FolderInfo struct {
	Name   string
	IsFile bool
}
