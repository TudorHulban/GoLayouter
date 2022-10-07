package main

type entry struct {
	folderInfo string
	indent     int
}

type entries []*entry

func NewEntries(content []string) *entries {
	var res entries

	for _, line := range content {
		res = append(res, convertToEntry(line))
	}

	return &res
}

func (e *entries) parse() []string {
	var res []string

	var stackFolders stack
	var stackIndents stack
	var stackPackages stack

	for ix, entry := range *e {
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
