package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	return emptyFile.Close()
}

func changeDirectory(path string) {
	os.Chdir(path)
}

// readFile is a helper reading file contents to a slice.
func readFile(filePath string) ([]string, error) {
	fileHandler, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer fileHandler.Close()

	var res []string

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil
}

// type osFunc func(path string) error

// var actions = make(map[string]osFunc)

func checkIfFileExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	return nil
}

func clearFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	return f.Close()
}

func writeInFile(message, path string) {
	fileHandler, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprintln(fileHandler, message)
	if err != nil {
		fmt.Println(err)
	}
}

func readByLine(fileName string) []string {
	var res []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		res = append(res, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		log.Panic(err)
	}

	return res
}
