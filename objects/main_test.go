package objects

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeofFile(t *testing.T) {
	got := []string{"! .", "# package main", "file.go", "folder"}
	want := []string{"path", "package", "file", "folder"}

	for i := range got {
		assert.Equal(t, typeofFile(got[i]), want[i], "verify the type of file")
		fmt.Println(got[i], "is a", want[i])
	}
}

func TestWriteToDisk(t *testing.T) {
	content, errRe := readFile(_pathInput)
	require.NoError(t, errRe)

	entries := parse(convertToEntries(content))

	for _, file := range entries {
		err := checkIfExist(file)
		assert.Equal(t, err, nil, "No match for file", file)
	}
}

func TestWriteToFile(t *testing.T) {
	content, errRe := readFile(_pathInput)
	require.NoError(t, errRe)

	clearFile(_pathOutput)
	errWr := writeToFile(_pathInput, _pathOutput)
	require.NoError(t, errWr)

	output, _ := readByLine(_pathOutput)
	entries := parse(convertToEntries(content))

	assert.Equal(t, output, entries, "should be equal")
}
