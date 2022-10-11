package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
	"github.com/TudorHulban/GoLayouter/objects"
)

const _pathInput = "test_cases/folder_c1"
const _pathOutput = "test_cases/folder_c1_results"

func TestTypeofFile(t *testing.T) {
	got := []string{"! .", "# package main", "file.go", "folder"}
	want := []string{"path", "package", "file", "folder"}

	for i := range got {
		assert.Equal(t, helpers.TypeofFile(got[i]), want[i], "verify the type of file")
		fmt.Println(got[i], "is a", want[i])
	}
}

func TestWriteToDisk(t *testing.T) {
	content, errRe := helpers.ReadFile(_pathInput)
	require.NoError(t, errRe)

	var e *objects.Entries

	entries := e.Parse(objects.NewEntries(content))

	for _, file := range entries {
		err := helpers.CheckIfExist(file)
		assert.Equal(t, err, nil, "No match for file", file)
	}
}

func TestWriteToFile(t *testing.T) {
	content, errRe := helpers.ReadFile(_pathInput)
	require.NoError(t, errRe)

	helpers.ClearFile(_pathOutput)
	errWr := helpers.WriteToFile(_pathInput, _pathOutput)
	require.NoError(t, errWr)

	var e *objects.Entries

	output, _ := helpers.ReadByLine(_pathOutput)
	entries := e.Parse(objects.NewEntries(content))

	assert.Equal(t, output, entries, "should be equal")
}
