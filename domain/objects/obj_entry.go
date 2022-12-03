package objects

import (
	"strings"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/app/helpers/stack"
)

type entry struct {
	folderInfo string
	indent     int
}

type Entries []*entry

const _defaultPackage = "package main"

func convertToEntry(lineOfText string) *entry {
	trimmed := strings.TrimLeft(lineOfText, " ")

	return &entry{
		folderInfo: trimmed,
		indent:     len(lineOfText) - len(trimmed),
	}
}

func NewEntries(content []string) *Entries {
	var res Entries

	for _, line := range content {
		res = append(res, convertToEntry(line))
	}

	return &res
}

func (e *Entries) Parse() []string {
	var res []string

	var stackFolders stack.Stack
	var stackIndents stack.Stack
	var stackPackages stack.Stack

	for ix, entry := range *e {
		if helpers.TypeofFile(entry.folderInfo) == "path" {
			stackFolders = nil
			stackIndents = nil
			stackPackages = nil

			if helpers.GetCommand(entry.folderInfo) != "." {
				stackFolders.Push(helpers.GetCommand(entry.folderInfo))
				res = append(res, stackFolders.String())
				stackIndents.Push(-1)

				continue
			}

			stackIndents.Push(entry.indent)

			continue
		}

		if helpers.TypeofFile(entry.folderInfo) == "pack" {
			stackPackages.Push(helpers.GetCommand(entry.folderInfo))

			continue
		}

		if helpers.TypeofFile(entry.folderInfo) == "file" {
			pack := stackPackages.Peek()

			if stackPackages.IsEmpty() {
				stackPackages.Push(_defaultPackage)

				pack = _defaultPackage
			}

			if stackPackages.Peek() == "t" {
				stackPackages.Pop()

				pack = stackPackages.Peek().(string)

				stackPackages.Push("t")
			}

			files := helpers.ConvertToFiles(entry.folderInfo, stackPackages.Peek().(string))

			for _, file := range files {
				file = file + "(" + pack.(string) + ")"
				line := stackFolders.String() + "/" + file
				res = append(res, line)
			}

			continue
		}

		if ix == 0 {
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(0)

			res = append(res, stackFolders.String())

			continue
		}

		if stackIndents.Peek().(int) < 0 {
			res = res[:len(res)-1]
		}

		if entry.indent > stackIndents.Peek().(int) {
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)

			res = append(res, stackFolders.String())

			continue
		}

		if entry.indent == stackIndents.Peek().(int) {
			stackFolders.Pop()
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)

			res = append(res, stackFolders.String())

			continue
		}

		for entry.indent < stackIndents.Peek().(int) && len(stackIndents) > 1 {
			stackFolders.Pop()
			stackIndents.Pop()

			if entry.indent == stackIndents.Peek().(int) {
				stackFolders.Pop()
				stackPackages.Pop()

				break
			}
		}

		stackFolders.Push(entry.folderInfo)
		stackIndents.Push(entry.indent)

		res = append(res, stackFolders.String())
	}

	return res
}
