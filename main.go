package main

import (
	"strings"
)

type entry struct {
	folderInfo string
	indent     int
}

const _filePath = "folders.txt"

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

func isTestFile(isPackage, line string) string {
	if isPackage == "t" {
		return line[:len(line)-3] + "_test.go"
	}

	return line
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

func getPackage(line string) string {
	return line[2:]
}

func convertToEntry(line string) *entry {
	trimmed := strings.TrimLeft(line, " ")

	return &entry{
		folderInfo: trimmed,
		indent:     len(line) - len(trimmed),
	}
}

func convertToEntries(content []string) []*entry {
	var res []*entry

	for _, line := range content {
		res = append(res, convertToEntry(line))
	}

	return res
}

func parse(entries []*entry) []string {
	var res []string

	var stackFolders stack
	var stackIndents stack
	var stackPackages stack

	for ix, entry := range entries {
		if typeofFile(entry.folderInfo) == "path" {
			stackFolders = nil
			stackIndents = nil
			stackPackages = nil

			res = append(res, getPackage(entry.folderInfo))

			stackIndents.push(getPackage(entry.folderInfo))
			changeDirectory(getPackage(entry.folderInfo))

			continue
		}

		if typeofFile(entry.folderInfo) == "pack" {
			stackPackages.push(getPackage(entry.folderInfo))

			continue
		}

		if typeofFile(entry.folderInfo) == "file" {
			pack := stackPackages.peek()
			files := lineParser(entry.folderInfo, pack.(string))

			for _, file := range files {
				line := stackFolders.String() + "/" + file
				res = append(res, line)

				//createFile(line)

				if pack != "t" {
					//	writeInFile(pack.(string), line)
				}
			}

			continue
		}

		if ix == 0 {
			stackFolders.push(entry.folderInfo)
			stackIndents.push(0)

			res = append(res, stackFolders.String())
			//createFolder(stackFolders.String())

			continue
		}

		if entry.indent > stackIndents.peek().(int) {
			stackFolders.push(entry.folderInfo)
			stackIndents.push(entry.indent)
			stackPackages.push("")

			res = append(res, stackFolders.String())
			//createFolder(stackFolders.String())

			continue
		}

		if entry.indent == stackIndents.peek().(int) {
			stackFolders.pop()
			stackFolders.push(entry.folderInfo)
			stackIndents.push(entry.indent)
			stackPackages.push("")

			res = append(res, stackFolders.String())
			//createFolder(stackFolders.String())

			continue
		}

		for entry.indent < stackIndents.peek().(int) && len(stackIndents) > 1 {
			if entry.indent == stackIndents.peek().(int) {
				stackFolders.pop()
				stackPackages.pop()

				break
			}

			stackFolders.pop()
			stackIndents.pop()
		}

		stackFolders.push(entry.folderInfo)
		stackIndents.push(entry.indent)

		res = append(res, stackFolders.String())
		//createFolder(stackFolders.String())
	}

	return res
}

const _pathInput = "test_cases/folder_c1"
const _pathOutput = "test_cases/folder_c1_results"

func main() {
}
