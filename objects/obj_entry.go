package objects

import (
	"strings"

	"github.com/TudorHulban/GoLayouter/helpers"
	s "github.com/TudorHulban/GoLayouter/stack"
)

type entry struct {
	folderInfo string
	indent     int
}

type Entries []*entry

func NewEntries(content []string) *Entries {
	var res Entries

	for _, line := range content {
		res = append(res, convertToEntry(line))
	}

	return &res
}

func convertToEntry(line string) *entry {
	trimmed := strings.TrimLeft(line, " ")

	return &entry{
		folderInfo: trimmed,
		indent:     len(line) - len(trimmed),
	}
}

func (e *Entries) Parse(*Entries) []string {
	var res []string

	var stackFolders s.Stack
	var stackIndents s.Stack
	var stackPackages s.Stack

	for ix, entry := range *e {
		if helpers.TypeofFile(entry.folderInfo) == "path" {
			stackFolders = nil
			stackIndents = nil
			stackPackages = nil

			res = append(res, helpers.GetPackage(entry.folderInfo))

			stackIndents.Push(helpers.GetPackage(entry.folderInfo))
			helpers.ChangeDirectory(helpers.GetPackage(entry.folderInfo))

			continue
		}

		if helpers.TypeofFile(entry.folderInfo) == "pack" {
			stackPackages.Push(helpers.GetPackage(entry.folderInfo))

			continue
		}

		if helpers.TypeofFile(entry.folderInfo) == "file" {
			pack := stackPackages.Peek()
			files := helpers.ConvertToFiles(entry.folderInfo, pack.(string))

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
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(0)

			res = append(res, stackFolders.String())
			//createFolder(stackFolders.String())

			continue
		}

		if entry.indent > stackIndents.Peek().(int) {
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)
			stackPackages.Push("")

			res = append(res, stackFolders.String())
			//createFolder(stackFolders.String())

			continue
		}

		if entry.indent == stackIndents.Peek().(int) {
			stackFolders.Pop()
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)
			stackPackages.Push("")

			res = append(res, stackFolders.String())
			//createFolder(stackFolders.String())

			continue
		}

		for entry.indent < stackIndents.Peek().(int) && len(stackIndents) > 1 {
			if entry.indent == stackIndents.Peek().(int) {
				stackFolders.Pop()
				stackPackages.Pop()

				break
			}

			stackFolders.Pop()
			stackIndents.Pop()
		}

		stackFolders.Push(entry.folderInfo)
		stackIndents.Push(entry.indent)

		res = append(res, stackFolders.String())
		//createFolder(stackFolders.String())
	}
	return res
}
