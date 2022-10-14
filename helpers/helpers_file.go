package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PathInput = "../test_cases/folder_c1"
const PathOutput = "../test_cases/folder_c1_results"

func TypeofFile(fileName string) string {
	if strings.Contains(fileName, "!") {
		return "path"
	}
	if strings.Contains(fileName, ".") {
		return "file"
	}
	if strings.Contains(fileName, "#") {
		return "pack"
	}

	return "folder"
}

// ReadFile  is a helper reading file contents to a slice.
func ReadFile(filePath string) ([]string, error) {
	fileHandler, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	err = fileHandler.Close()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ReadByLine readFile is a helper reading file contents line by line.
func ReadByLine(fileName string) ([]string, error) {
	var res []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		res = append(res, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		return nil, err
	}

	return res, nil
}

func CheckIfExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	return nil
}

func ChangeDirectory(path string) error {
	return os.Chdir(path)
}

func ClearFile(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func WriteLineInFile(message, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, message)
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
