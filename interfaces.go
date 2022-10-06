package main

type IWritter interface {
	writeToDisk() error
}

// TODO: move to another file
func writeToFile(input, output string) error {
	content, errRe := readFile(input)
	if errRe != nil {
		return errRe
	}

	entries := parse(convertToEntries(content))

	for _, file := range entries {
		writeInFile(file, output)
	}

	return nil
}
