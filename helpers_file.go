package main

import (
	"bufio"
	"os"
)

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
