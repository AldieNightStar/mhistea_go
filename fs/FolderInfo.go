package fs

type FolderInfoImp struct {
	name   string
	isFile bool
}

func NewFolderInfo(name string, isFile bool) FolderInfo {
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
