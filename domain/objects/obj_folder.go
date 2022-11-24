package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/interfaces"
)

type Folder struct {
	Path string
}

var _ interfaces.IFileOperations = &Folder{}

func (f *Folder) SetPath(path string) {
	f.Path = path
}

func (f Folder) DeletePath() error {
	return RemoveFile(f.GetPath())
}

func (f Folder) CheckIfPathExists() error {
	return helpers.CheckIfPathExists(f.GetPath())
}

func (f Folder) GetPath() string {
	return f.Path
}

func (f Folder) WriteToDisk() error {
	err := os.Mkdir(f.Path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
