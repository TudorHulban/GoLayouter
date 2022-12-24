package item

import (
	"github.com/TudorHulban/GoLayouter/app/helpers/stack"
	helpers "github.com/TudorHulban/GoLayouter/app/helpers/utils"
	interfaces "github.com/TudorHulban/GoLayouter/domain/intefaces"
	e "github.com/TudorHulban/GoLayouter/domain/objects/entry"
)

type Item struct {
	//object path
	path interfaces.IFileOperations

	kind string
}

func (e *e.Entries) Parse2() []Item {
	var res []Item

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

				res = append(res, Item{
					path: &Folder{
						Path: stackFolders.String(),
					},
					kind: helpers.KindofFile(entry.folderInfo),
				})

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

			if stackPackages.Peek() == "t" || stackPackages.Peek() == "tt" {
				stackPackages.Pop()

				pack = stackPackages.Peek().(string)

				stackPackages.Push("t")
			}

			files := helpers.ConvertToFiles(entry.folderInfo, stackPackages.Peek().(string))

			for _, file := range files {
				file = file + "(" + pack.(string) + ")"
				line := stackFolders.String() + "/" + file

				res = append(res, Item{
					path: &Folder{
						Path: line,
					},
					kind: helpers.KindofFile(line),
				})

			}

			continue
		}

		if ix == 0 {
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(0)

			res = append(res, Item{
				path: &Folder{
					Path: stackFolders.String(),
				},
				kind: helpers.KindofFile(entry.folderInfo),
			})

			continue
		}

		if stackIndents.Peek().(int) < 0 {
			res = res[:len(res)-1]
		}

		if entry.indent > stackIndents.Peek().(int) {
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)

			res = append(res, Item{
				path: &Folder{
					Path: stackFolders.String(),
				},
				kind: helpers.KindofFile(entry.folderInfo),
			})

			continue
		}

		if entry.indent == stackIndents.Peek().(int) {
			stackFolders.Pop()
			stackFolders.Push(entry.folderInfo)
			stackIndents.Push(entry.indent)

			res = append(res, Item{
				path: &Folder{
					Path: stackFolders.String(),
				},
				kind: helpers.KindofFile(entry.folderInfo),
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

		res = append(res, Item{
			path: &Folder{
				Path: stackFolders.String(),
			},
			kind: helpers.KindofFile(entry.folderInfo),
		})
	}

	return res
}
