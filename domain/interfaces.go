package domain

type IFileOperations interface {
	WriteToDisk() error
	CheckIfPathExists() error
	ChangeDirectory(newPath string) error
	DeletePath() error
	GetPath() string
}
