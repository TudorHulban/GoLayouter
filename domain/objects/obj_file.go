package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
)

type File struct {
	Path    string
	Content string
}

func (f *File) SetPath(path string) {
	(*f).Path = path
}

func (f *File) SetContent(content string) {
	f.Content = content
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
