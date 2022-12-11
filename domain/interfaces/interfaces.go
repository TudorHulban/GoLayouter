package interfaces

type IFileOperations interface {
	WriteToDisk() error

	ChangeDirectory(newPath string) error
	CheckIfPathExists() error

	DeletePath() error
}
