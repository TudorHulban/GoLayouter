package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
)

type Folder struct {
	Path string
}

func (f *Folder) SetPath(path string) {
	(*f).Path = path
}

func (f Folder) DeletePath() error {
	return RemoveFile(f.GetPath())
}

func (f Folder) CheckIfExists() error {
	return helpers.CheckIfPathExists(f.GetPath())
}

func (f Folder) GetPath() string {
	return f.Path
}

//var _ interfaces.IWritter = &Folder{}

func (f Folder) WriteToDisk() error {
	err := os.Mkdir(f.Path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
