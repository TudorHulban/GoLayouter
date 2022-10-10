package objects

type folder struct {
	path string
}

func (f folder) writeToDisk() error {
	return createFolder(f.path)
}

func convertToFolder(line string) *folder {
	return &folder{
		path: line,
	}
}
