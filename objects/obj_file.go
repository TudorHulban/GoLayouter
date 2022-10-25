package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/helpers"
)

type File struct {
	Path    string
	Content string
}

//var _ interfaces.IWritter = &File{}

func (f File) CreateFile() error {
	emptyFile, err := os.Create(f.Path)
	if err != nil {
		return err
	}

	err = helpers.WriteTextInFile(f.Content, f.Path)
	if err != nil {
		return err
	}

	return emptyFile.Close()
}

func RemoveFile(path string) error {
	return os.Remove(helpers.RemovePackageName(path))
}
