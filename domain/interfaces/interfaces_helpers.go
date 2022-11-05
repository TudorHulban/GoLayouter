package interfaces

import (
	"bufio"
	"os"
)

const _pathInput = "../test_cases/folder_c6"

// to evoid imported cycle
func IRWritterReadFile(filePath string) ([]string, error) {
	fileHandler, errOp := os.Open(filePath)
	if errOp != nil {
		return nil, errOp
	}

	var errClo error
	defer func() {
		errClo = fileHandler.Close()
	}()

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, errClo
}

func WriteToDisk(paths []IWritter) error {
	for _, path := range paths {
		err := path.WriteToDisk()
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckPathsExists(paths []IWritter) error {
	for _, path := range paths {
		errCheck := path.CheckIfExists()
		if errCheck != nil {
			return errCheck
		}
	}

	return nil
}

func DeletePaths(paths []IWritter) error {
	for index := len(paths) - 1; index >= 0; index-- {
		err := paths[index].DeletePath()
		if err != nil {
			return err
		}
	}

	return nil
}
