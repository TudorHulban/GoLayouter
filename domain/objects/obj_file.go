package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/interfaces"
)

type File struct {
	Path    string
	Content string
}

var _ interfaces.IFileOperations = &File{}

func (f File) SetPath(path string) {
	f.Path = path
}

func (f *File) SetContent(content string) {
	f.Content = content
}

func (f File) GetPath() string {
	return f.Path
}

func (f File) CheckIfPathExists() error {
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
