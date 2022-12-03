package interfaces

type IFileOperations interface {
	WriteToDisk() error

	CheckIfPathExists() error
	DeletePath() error
}
