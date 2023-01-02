package domain

import "os"

type IFileOperations interface {
	WriteToDisk() (*os.File, error)
	CheckIfPathExists() error
	ChangeDirectory(newPath string) error
	DeletePath() error
}
