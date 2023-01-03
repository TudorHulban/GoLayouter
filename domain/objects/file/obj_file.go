package file

import (
	"io"
	"os"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/domain"
)

type File struct {
	// extracted from initial file the path where the file will be created
	Path              string
	GolangPackageName string
}

var _ domain.IFileOperations = &File{}
var _ io.Writer = File{}

func (f File) GetPath() string {
	return f.Path
}

// Write opens the file path and write the content parsed with "Write"
// method from io package. If file does not exists
// it will be created with "Create" method from os package.
func (f File) Write(content []byte) (int, error) {
	if helpers.CheckIfPathExists(f.Path) != nil {
		path, errCreate := os.Create(f.Path)
		if errCreate != nil {
			return 0, errCreate
		}

		return path.Write(content)
	}

	path, errOpen := os.Open(f.Path)
	if errOpen != nil {
		return 0, errOpen
	}

	return path.Write(content)
}

func (f File) CheckIfPathExists() error {
	return helpers.CheckIfPathExists(f.Path)
}

func (f File) DeletePath() error {
	return os.Remove(f.Path)
}

func (f File) WriteToDisk() error {
	_, errCreate := os.Create(f.Path)
	return errCreate
}

func (f *File) ChangeDirectory(newPath string) error {
	errPathExists := helpers.CheckIfPathExists(newPath)
	if errPathExists != nil {
		return errPathExists
	}

	f.Path = newPath + "/" + f.Path

	return nil
}
