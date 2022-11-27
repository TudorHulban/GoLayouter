package service

import (
	"errors"
	"path"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/interfaces"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

type Service struct {
	paths []interfaces.IFileOperations
}

func NewService(content []string) (*Service, error) {
	if len(content) == 0 {
		return nil, errors.New("parsed content is empty")
	}

	var res []interfaces.IFileOperations

	for _, line := range content {
		_, fileName := path.Split(line)
		if helpers.TypeofFile(fileName) == "file" {
			packageName := helpers.ParsePackage(fileName)
			path := helpers.RemovePackageName(line)

			res = append(res, &objects.File{
				Path:    path,
				Content: packageName,
			})

			continue
		}
		_, folderName := path.Split(line)
		if helpers.TypeofFile(folderName) == "folder" {
			res = append(res, &objects.Folder{Path: line})
		}
	}

	return &Service{
		paths: res,
	}, nil
}

func (serv *Service) WriteToDisk() error {
	for _, path := range serv.paths {
		if err := path.WriteToDisk(); err != nil {
			return err
		}
	}

	return nil
}

func (Service) ConvertToIFileOperations(content []string) []interfaces.IFileOperations {
	res := make([]interfaces.IFileOperations, len(content), len(content))

	for ix, line := range content {
		_, fileName := path.Split(line)
		if helpers.TypeofFile(fileName) == "file" {
			packageName := helpers.ParsePackage(fileName)
			path := helpers.RemovePackageName(line)

			res[ix] = &objects.File{
				Path:    path,
				Content: packageName,
			}

			continue
		}
		_, folderName := path.Split(line)
		if helpers.TypeofFile(folderName) == "folder" {
			res[ix] = &objects.File{
				Path: line,
			}
		}
	}

	return res
}
