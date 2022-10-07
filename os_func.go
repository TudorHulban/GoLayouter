package main

import (
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

func writeLineInFile(message, fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprintln(f, message)
	if err != nil {
		fmt.Println(err)
	}
}

func writeToFile(input, output string) error {
	content, errRe := readFile(input)
	if errRe != nil {
		return errRe
	}

	entries := parse(convertToEntries(content))

	for _, file := range entries {
		writeLineInFile(file, output)
	}

	return nil
}
