package objects

import (
	"os"
)

type Folder struct {
	Path string
}

//var _ interfaces.IWritter = &Folder{}

func (f Folder) WriteToDisk() error {
	err := os.Mkdir(f.Path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
