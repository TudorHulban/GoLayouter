package main

type folder struct {
	path string
}

func (f folder) writeToDisk() error {
	return createFolder(f.path)
}
