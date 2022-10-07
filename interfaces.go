package main

type IWritter interface {
	writeToDisk() error
<<<<<<< HEAD
=======
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
>>>>>>> ceaffc7bb98a80cc5b04bd9a404c0b5f9181c1f9
}
