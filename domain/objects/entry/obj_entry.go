package entry

import (
	"strings"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/app/stack"
	"github.com/TudorHulban/GoLayouter/domain/objects/file"
	"github.com/TudorHulban/GoLayouter/domain/objects/folder"
	"github.com/TudorHulban/GoLayouter/domain/objects/item"
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

func (e *Entries) ParseToItems() []item.Item {
	var res []item.Item

	var first bool

	var stackFolders stack.Stack
	var stackIndents stack.Stack
	var stackPackages stack.Stack

	for _, entry := range *e {
		if helpers.TypeofFile(entry.folderInfo) == "path" {
			stackFolders = nil
			stackIndents = nil
			stackPackages = nil

			if helpers.GetCommand(entry.folderInfo) != "." {
				stackFolders.Push(helpers.GetCommand(entry.folderInfo))
				stackIndents.Push(-1)

				res = append(res, item.Item{
					ObjectPath: &folder.Folder{
						Path: stackFolders.String(),
					},
					Kind: helpers.KindofFile(entry.folderInfo),
				})

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
			packageName := stackPackages.Peek()

			if stackPackages.IsEmpty() {
				stackPackages.Push(_defaultPackage)

				packageName = _defaultPackage
			}

			if stackPackages.Peek() == "t" || stackPackages.Peek() == "tt" {
				testPackage := stackPackages.Pop()
				packageName = stackPackages.Peek().(string)

				stackPackages.Push(testPackage)
			}

			files := helpers.LineToFiles(entry.folderInfo, stackPackages.Peek().(string))
			for _, fileName := range files {
				res = append(res, item.Item{
					ObjectPath: &file.File{
						Path:              stackFolders.String() + "/" + fileName,
						GolangPackageName: packageName.(string),
					},
					Kind: helpers.KindofFile(fileName),
				})
			}

			continue
		}

		if (stackIndents != nil) && (stackIndents.Peek().(int) < 0) {
			res = res[:len(res)-1]
		}

		if !first {
			first = true

			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)

			folder := &folder.Folder{
				Path: stackFolders.String(),
			}
			res = append(res, item.Item{
				ObjectPath: folder,
				Kind:       helpers.KindofFile(folder.Path),
			})

			continue
		}

		if entry.indent > stackIndents.Peek().(int) {
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)

			res = append(res, item.Item{
				ObjectPath: &folder.Folder{
					Path: stackFolders.String(),
				},
				Kind: helpers.KindofFile(entry.folderInfo),
			})

			continue
		}

		if entry.indent == stackIndents.Peek().(int) {
			stackFolders.Pop()
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)

			res = append(res, item.Item{
				ObjectPath: &folder.Folder{
					Path: stackFolders.String(),
				},
				Kind: helpers.KindofFile(entry.folderInfo),
			})

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

		res = append(res, item.Item{
			ObjectPath: &folder.Folder{
				Path: stackFolders.String(),
			},
			Kind: helpers.KindofFile(entry.folderInfo),
		})

	}

	return res
}
