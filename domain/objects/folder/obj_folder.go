package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/app/utils/helpers"
	"github.com/TudorHulban/GoLayouter/domain/interfaces"
)

type Folder struct {
	Path string // extracted from initial file the path where the folder will be created
}

var _ interfaces.IFileOperations = &Folder{}

func (f Folder) DeletePath() error {
	return os.Remove(f.Path)
}

func (f Folder) CheckIfPathExists() error {
	return helpers.CheckIfPathExists(f.Path)
}

func (f Folder) WriteToDisk() error {
	return helpers.CreateFolder(f.Path)
}

func (f *Folder) ChangeDirectory(newPath string) error {
	errPathExists := helpers.CheckIfPathExists(newPath)
	if errPathExists != nil {
		return errPathExists
	}

	f.Path = newPath + "/" + f.Path

	return nil
}
