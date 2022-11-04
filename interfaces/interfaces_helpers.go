package interfaces

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/objects"
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

func TestConvertToIWritter(t *testing.T) {
	content, errRead := IRWritterReadFile(_pathInput)
	require.NoError(t, errRead)

	e := objects.NewEntries(content)
	entries := e.Parse()
	writter := objects.ConvertToIWritter(entries)

	for _, element := range writter {
		log.Print(element)
	}
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
