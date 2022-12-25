package domain

// TODO : remove interface/interfaces
type IFileOperations interface {
	WriteToDisk() error
	CheckIfPathExists() error
	ChangeDirectory(newPath string) error
	DeletePath() error
}
