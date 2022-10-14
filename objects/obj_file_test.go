package objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
)

func TestCreateFile(t *testing.T) {
	fileName := "file.go"

	errCreate := CreateFile(fileName)
	require.NoError(t, errCreate)

	errCheck := helpers.CheckIfExist(fileName)
	require.NoError(t, errCheck)

	assert.Equal(t, errCheck, nil, "No error should be returned while checking")
}

func TestRemoveFile(t *testing.T) {
	fileName := "file.go"

	errCreate := CreateFile(fileName)
	require.NoError(t, errCreate)

	errRemoveFile := RemoveFile(fileName)
	require.NoError(t, errRemoveFile)

	fileExist := helpers.CheckIfExist(fileName)

	assert.NotEqual(t, fileExist, nil, "no nil should be returned while checking")
}

func TestGetFile(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		output      string
	}{
		{"1 file", "folder-root1", "folder-root1"},
		{"2 files", "folder-root1/subfolder1", "subfolder1"},
		{"3 files", "folder-root1/subfolder1/subsubfolder1", "subsubfolder1"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.output, GetFile(tc.input))
		})
	}

}

func TestWriteObjectsToFile(t *testing.T) {
	content, errRe := helpers.ReadByLine(helpers.PathInput)
	require.NoError(t, errRe)

	errClearFile := helpers.ClearFile(helpers.PathOutput)
	require.NoError(t, errClearFile)

	errWr := WriteToFile(helpers.PathInput, helpers.PathOutput)
	require.NoError(t, errWr)

	e := NewEntries(content)
	entries := e.Parse()

	output, errRead := helpers.ReadByLine(helpers.PathOutput)
	require.NoError(t, errRead)

	assert.Equal(t, output, entries, "should be equal")
}
