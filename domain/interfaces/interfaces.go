package interfaces

// TODO : remove interface/interfaces
type IFileOperations interface {
	WriteToDisk() error

	ChangeDirectory(newPath string) error
	CheckIfPathExists() error

	DeletePath() error
}
