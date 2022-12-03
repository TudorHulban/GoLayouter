package objects

import (
	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/interfaces"
)

type Folder struct {
	Path string // extracted from initial file the path where the folder will be created
}

var _ interfaces.IFileOperations = &Folder{}

func (f Folder) GetPath() string {
	return f.Path
}

func (f Folder) DeletePath() error {
	return RemoveFile(f.GetPath())
}

func (f Folder) CheckIfPathExists() error {
	return helpers.CheckIfPathExists(f.GetPath())
}

func (f Folder) WriteToDisk() error {
	return helpers.CreateFolder(f.Path)
}
