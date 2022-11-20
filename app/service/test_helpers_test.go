package service

import "github.com/TudorHulban/GoLayouter/domain/interfaces"

func CheckPathsExists(paths []interfaces.IFileOperations) error {
	for _, path := range paths {
		errCheck := path.CheckIfPathExists()
		if errCheck != nil {
			return errCheck
		}
	}

	return nil
}

func DeletePaths(paths []interfaces.IFileOperations) error {
	for index := len(paths) - 1; index >= 0; index-- {
		err := paths[index].DeletePath()
		if err != nil {
			return err
		}
	}

	return nil
}
