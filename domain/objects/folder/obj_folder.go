package folder

import (
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/domain"
)

type Folder struct {
	Path string // extracted from initial file the path where the folder will be created
}

var _ domain.IFileOperations = &Folder{}

func (f Folder) DeletePath() error {
	return os.Remove(f.Path)
}

func (f Folder) CheckIfPathExists() error {
	return helpers.CheckIfPathExists(f.Path)
}

func (f Folder) WriteToDisk() (*os.File, error) {
	errCreate := helpers.CreateFolder(f.Path)
	if errCreate != nil {
		return nil, errCreate
	}

	file, errOpenFile := os.Open(f.Path)
	if errOpenFile != nil {
		return nil, errOpenFile
	}

	return file, nil
}

func (f *Folder) ChangeDirectory(newPath string) error {
	errPathExists := helpers.CheckIfPathExists(newPath)
	if errPathExists != nil {
		return errPathExists
	}

	(*f).Path = newPath + "/" + f.Path

	return nil
}
