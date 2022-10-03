package main

import (
	"os"
)

type osFunc func(path string) error

var actions = make(map[string]osFunc)

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

	emptyFile.Close()

	return nil
}

func checkIfExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	return nil
}

func changeDirectory(path string) {
	os.Chdir(path)
}
