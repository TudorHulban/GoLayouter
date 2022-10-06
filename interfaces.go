package main

type IWritter interface {
	writeToDisk() error
	writeToFile() error
}

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
