package common

type FileReader interface {
	// Read file
	ReadFile(name string) []byte
}

type FileWriter interface {
	// Write File (Rewrite if exists)
	WriteFile(name string, data []byte) bool
}

type FileLister interface {
	// List of files and folders
	List() []FolderInfo
}

type Folder interface {
	FileReader
	FileWriter
	FileLister

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

type FolderInfo interface {
	Name() string
	IsFile() bool
}
