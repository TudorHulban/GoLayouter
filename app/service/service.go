package service

import (
	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/interfaces"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

type service struct{}

func (service) WriteToDisk(paths []interfaces.IWritter) error {
	for _, path := range paths {
		err := path.WriteToDisk()
		if err != nil {
			return err
		}
	}

	return nil
}

func (service) ConvertToIWritter(content []string) []interfaces.IWritter {
	res := make([]interfaces.IWritter, len(content), len(content))

	for _, line := range content {
		if helpers.TypeofFile(helpers.GetFileName(line)) == "file" {
			packageName := helpers.ParsePackage(helpers.GetFileName(line))
			path := helpers.RemovePackageName(line)

			res = append(res, &objects.File{
				Path:    path,
				Content: packageName,
			})

			continue
		}

		if helpers.TypeofFile(helpers.GetFileName(line)) == "folder" {
			res = append(res, &objects.Folder{Path: line})
		}
	}

	return res
}
