package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetCommand(line string) string {
	return line[2:]
}

func IsTestFile(text string) bool {
	return text == "t"
}

func CreateGolangTestFile(text string) string {
	return text[:len(text)-3] + "_test.go"
}

func ConvertToFiles(text, packageName string) []string {
	var res []string
	files := strings.Split(text, " ")

	for _, file := range files {
		fileTrimmed := strings.TrimLeft(file, " ")

		if fileTrimmed != "" {
			if IsTestFile(packageName) {
				res = append(res, fileTrimmed, CreateGolangTestFile(fileTrimmed))

				continue
			}

			res = append(res, fileTrimmed)
		}
	}

	return res
}

func ParsePackage(text string) string {
	var start, stop int

	for ix, character := range text {
		if character == '(' {
			start = ix + 1

			continue
		}

		if character == ')' {
			stop = ix
		}
	}

	return text[start:stop]
}

func GetFileName(path string) string {
	var res string
	var found bool

	for ix, character := range path {
		if character == '/' {
			res = path[ix+1:]
			found = true
		}
	}

	if !found {
		res = path
	}

	return res
}

func RemovePackageName(text string) string {
	var stop int

	for ix, character := range text {
		if character == '(' {
			stop = ix

			break
		}
	}

	if stop == 0 {
		return text
	}

	return text[:stop]
}

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

func CheckIfPathExists(path string) error {
	_, errStat := os.Stat(path)
	if errStat == nil {
		return nil
	}

	if errors.Is(errStat, os.ErrNotExist) {
		return errStat
	}

	return fmt.Errorf("os error: %w", errStat)
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
