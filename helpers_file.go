package main

import (
	"bufio"
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
