package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	return err
}

func ClearFile(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer func() {
		errClose := f.Close()
		if errClose != nil {
			err = errClose
		}
	}()

	return err
}

func WriteTextInFile(text, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_, err = fmt.Fprintln(f, text)

		errClose := f.Close()
		if errClose != nil {
			err = errClose
		}
	}()

	return err
}
