package objects

import (
	"github.com/TudorHulban/GoLayouter/helpers"
)

type file struct {
	path    string
	content string
}

func (f file) writeToDisk() error {
	return helpers.CreateFile(f.path)
}

func convertToFile(line, packageName string) *file {
	var f file

	f.path = helpers.IsTestFile(packageName, line)

	if packageName == "" {
		f.content = "package main"
	} else {
		f.content = helpers.GetPackage(packageName)
	}

	return &file{
		path:    f.path,
		content: f.content,
	}
}
