package main

type file struct {
	path    string
	content string
}

func (f file) writeToDisk() error {
	return createFile(f.path)
}
