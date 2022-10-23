package objects

import (
	"os"
	"strings"

	"github.com/TudorHulban/GoLayouter/helpers"
	"github.com/TudorHulban/GoLayouter/interfaces"
)

type File struct {
	Path        string
	HasTestFile bool
}

var _ interfaces.IWritter = &File{}

func (f *File) CreateTestFile() error {
	// TODO: identify test file name and create a path
	// TODO: create file as per above path

	return nil
}

func getCommand(line string) string {
	return line[2:]
}

func isTestFile(text string) bool {
	return text == "t"
}

func createGolangTestFile(text string) string {
	return text[:len(text)-3] + "_test.go"
}

func convertToFiles(text, packageName string) []string {
	var res []string
	files := strings.Split(text, " ")

	for _, file := range files {
		fileTrimmed := strings.TrimLeft(file, " ")

		if fileTrimmed != "" {
			if isTestFile(packageName) {
				res = append(res, fileTrimmed, createTestFile(fileTrimmed))

				continue
			}

			res = append(res, fileTrimmed)
		}
	}

	return res
}

func GetFile(fileName string) string {
	var res string
	var found bool

	for ix, character := range fileName {
		if character == '/' {
			res = fileName[ix+1:]
			found = true
		}
	}

	if !found {
		res = fileName
	}

	return res
}

func RemoveFile(fileName string) error {
	return os.Remove(helpers.RemovePackageName(fileName))
}

func WriteToFile(entries []string, output string) error {
	for _, file := range entries {
		if err := helpers.WriteTextInFile(file, output); err != nil {
			return err
		}
	}

	return nil
}

func CreateFile(path string) error {
	emptyFile, err := os.Create(path)
	if err != nil {
		return err
	}

	return emptyFile.Close()
}
