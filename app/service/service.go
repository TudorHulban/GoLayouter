package service

import (
	"log"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/interfaces"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

type Service struct {
	paths []interfaces.IFileOperations
}

func NewService(content []string) *Service {
	var res []interfaces.IFileOperations

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

	return &Service{
		paths: res,
	}
}

func (serv *Service) WriteToDisk() error {
	for _, path := range serv.paths {
		log.Print(path)
		err := path.WriteToDisk()
		if err != nil {
			return err
		}
	}

	return nil
}

func (Service) ConvertToIFileOperations(content []string) []interfaces.IFileOperations {
	var res []interfaces.IFileOperations

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
