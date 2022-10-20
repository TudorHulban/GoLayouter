package objects

import (
	"os"
	"strings"

	"github.com/TudorHulban/GoLayouter/helpers"
)

func getPackage(line string) string {
	return line[2:]
}

func isTestFile(packageName string) bool {
	return packageName == "t"
}

func createTestFile(line string) string {
	return line[:len(line)-3] + "_test.go"
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
	found := false

	for i, character := range fileName {
		if character == '/' {
			res = fileName[i+1:]
			found = true
		}
	}

	if !found {
		res = fileName
	}

	return res
}

func RemoveFile(fileName string) error {
	errRm := os.Remove(helpers.RemovePackageName(fileName))
	if errRm != nil {
		return errRm
	}

	return nil
}

func WriteToFile(entries []string, output string) error {
	for _, file := range entries {
		err := helpers.WriteTextInFile(file, output)
		if err != nil {
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

	err = emptyFile.Close()
	if err != nil {
		return err
	}

	return nil
}
