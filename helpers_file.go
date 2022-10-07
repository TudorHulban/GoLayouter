package main

import (
	"bufio"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> ceaffc7bb98a80cc5b04bd9a404c0b5f9181c1f9
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

<<<<<<< HEAD
// readFile is a helper reading file contents line by line.
func readByLine(fileName string) ([]string, error) {
=======
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
>>>>>>> ceaffc7bb98a80cc5b04bd9a404c0b5f9181c1f9
	var res []string

	file, err := os.Open(fileName)
	if err != nil {
<<<<<<< HEAD
		return nil, err
=======
		log.Panic(err)
>>>>>>> ceaffc7bb98a80cc5b04bd9a404c0b5f9181c1f9
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		res = append(res, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		log.Panic(err)
	}

<<<<<<< HEAD
	return res, nil
=======
	return res
>>>>>>> ceaffc7bb98a80cc5b04bd9a404c0b5f9181c1f9
}
