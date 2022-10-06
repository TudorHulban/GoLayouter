package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteToDisk(t *testing.T) {
	content, errRe := readFile(_pathInput)
	require.NoError(t, errRe)

	entries := parse(convertToEntries(content))

	for _, file := range entries {
		err := checkIfFileExists(file)
		writeInFile(file, _pathOutput)
		assert.Equal(t, err, nil, "No match for file", file)
	}
}

func TestWriteToFile(t *testing.T) {
	content, errRe := readFile(_pathInput)
	require.NoError(t, errRe)

	clearFile(_pathOutput)
	writeToFile(_pathInput, _pathOutput)

	output := readByLine(_pathOutput)
	entries := parse(convertToEntries(content))

	assert.Equal(t, output, entries, "shoud be equal")
}
