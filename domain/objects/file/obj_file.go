package file

import (
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/domain"
)

type File struct {
	Path    string // extracted from initial file the path where the file will be created
	Content string
}

var _ domain.IFileOperations = &File{}

func (f File) CheckIfPathExists() error {
	return helpers.CheckIfPathExists(f.Path)
}

func (f File) DeletePath() error {
	return os.Remove(f.Path)
}

func (f File) WriteToDisk() error {
	var emptyFile *os.File
	if helpers.CheckIfPathExists(f.Path) != nil {
		var errCreate error

		emptyFile, errCreate = os.Create(f.Path)
		if errCreate != nil {
			return errCreate
		}

		errWrite := helpers.WriteTextInFile(f.Content, f.Path)
		if errWrite != nil {
			return errWrite
		}

		return emptyFile.Close()
	}

	return nil
}

func (f *File) ChangeDirectory(newPath string) error {
	errPathExists := helpers.CheckIfPathExists(newPath)
	if errPathExists != nil {
		return errPathExists
	}

	f.Path = newPath + "/" + f.Path

	return nil
}
