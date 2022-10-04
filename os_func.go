package main

import (
	"io/ioutil"
	"log"
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

func writeInFile(message, fileName string) {
	err := ioutil.WriteFile(fileName, []byte(message), 0644)
	if err != nil {
		log.Panic(err)
	}
}
