package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/helpers"
)

type File struct {
	Path    string
	Content string
}

func (f File) GetPath() string {
	return f.Path
}

func (f File) CheckIfExists() error {
	return helpers.CheckIfPathExists(f.GetPath())
}

func (f File) DeletePath() error {
	return RemoveFile(f.GetPath())
}

func (f File) WriteToDisk() error {
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
