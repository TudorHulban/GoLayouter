package objects

import "github.com/TudorHulban/GoLayouter/helpers"

type folder struct {
	path string
}

func (f folder) writeToDisk() error {
	return helpers.CreateFolder(f.path)
}

func convertToFolder(line string) *folder {
	return &folder{
		path: line,
	}
}
