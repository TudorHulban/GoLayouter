package objects

type file struct {
	path    string
	content string
}

func (f file) writeToDisk() error {
	return createFile(f.path)
}

func convertToFile(line, packageName string) *file {
	var f file

	f.path = isTestFile(packageName, line)

	if packageName == "" {
		f.content = "package main"
	} else {
		f.content = getPackage(packageName)
	}

	return &file{
		path:    f.path,
		content: f.content,
	}
}
