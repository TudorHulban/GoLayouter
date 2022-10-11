package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TudorHulban/GoLayouter/objects"
)

const _filePath = "folders.txt"

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

func IsTestFile(packageName, line string) string {
	if packageName == "t" {
		return line[:len(line)-3] + "_test.go"
	}

	return line
}

func GetPackage(line string) string {
	return line[2:]
}

func ConvertToFiles(text, packageName string) []string {
	var res []string
	files := strings.Split(text, " ")

	for _, file := range files {
		fileTrimmed := strings.TrimLeft(file, " ")

		if fileTrimmed != "" {
			res = append(res, IsTestFile(packageName, fileTrimmed), fileTrimmed)
		}
	}

	return res
}

// readFile is a helper reading file contents to a slice.
func ReadFile(filePath string) ([]string, error) {
	fileHandler, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	err = fileHandler.Close()
	if err != nil {
		return nil, err
	}

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil
}

// readFile is a helper reading file contents line by line.
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
		log.Panic(err)
	}

	return res, nil
}

func CreateFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
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

func CheckIfExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	return nil
}

func ChangeDirectory(path string) {
	os.Chdir(path)
}

func ClearFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func writeLineInFile(message, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, message)
	if err != nil {
		return err
	}

	return nil
}

func WriteToFile(input, output string) error {
	content, errRe := ReadFile(input)
	if errRe != nil {
		return errRe
	}

	var e *(objects.Entries)

	entries := e.Parse(objects.NewEntries(content))

	for _, file := range entries {
		err := writeLineInFile(file, output)
		if err != nil {
			return err
		}
	}

	return nil
}
