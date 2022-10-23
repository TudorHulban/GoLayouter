package objects

import (
	"os"

	"github.com/TudorHulban/GoLayouter/interfaces"
)

type Folder struct {
	Path string
}

var _ interfaces.IWritter = &Folder{}

func CreateFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

//func convertToFolder(line string) *folder {
//	return &folder{
//		path: line,
//	}
//}
