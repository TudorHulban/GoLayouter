package main

import (
	"strings"
)

const _filePath = "folders.txt"
const _pathInput = "test_cases/folder_c1"
const _pathOutput = "test_cases/folder_c1_results"

func typeofFile(fileName string) string {
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

func isTestFile(packageName, line string) string {
	if packageName == "t" {
		return line[:len(line)-3] + "_test.go"
	}

	return line
}

func getPackage(line string) string {
	return line[2:]
}

func lineParser(line, packageName string) []string {
	var res []string
	files := strings.Split(line, " ")

	for _, file := range files {
		fileTrimmed := strings.TrimLeft(file, " ")

		if fileTrimmed != "" {
			res = append(res, isTestFile(packageName, fileTrimmed), fileTrimmed)
		}
	}

	return res
}


func main() {
}
