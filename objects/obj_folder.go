package objects

import (
	"os"
)

//type folder struct {
//	path string
//}

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
