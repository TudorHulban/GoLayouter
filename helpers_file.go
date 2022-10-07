package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// readFile is a helper reading file contents to a slice.
func readFile(filePath string) ([]string, error) {
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
func readByLine(fileName string) ([]string, error) {
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

func createFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func createFile(path string) error {
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

func checkIfExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	return nil
}

//func changeDirectory(path string) {
//	os.Chdir(path)
//}

func clearFile(fileName string) {
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

func writeToFile(input, output string) error {
	content, errRe := readFile(input)
	if errRe != nil {
		return errRe
	}

	entries := parse(convertToEntries(content))

	for _, file := range entries {
		err := writeLineInFile(file, output)
		if err != nil {
			return err
		}
	}

	return nil
}
