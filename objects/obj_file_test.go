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
	testCases := []struct {
		description string
		fileInput   string
		fileOutput  string
	}{
		{"2 levels", "../test_cases/folder_c1", "../test_cases/folder_c1_results"},
		{"3 levels", "../test_cases/folder_c2", "../test_cases/folder_c2_results"},
		{"3 levels with going back", "../test_cases/folder_c3", "../test_cases/folder_c3_results"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRe := helpers.ReadByLine(tc.fileInput)
			require.NoError(t, errRe)

			errClearFile := helpers.ClearFile(tc.fileOutput)
			require.NoError(t, errClearFile)

			e := NewEntries(content)
			entries := e.Parse()

			errWr := WriteToFile(entries, tc.fileOutput)
			require.NoError(t, errWr)

			output, errRead := helpers.ReadByLine(tc.fileOutput)
			require.NoError(t, errRead)

			assert.Equal(t, output, entries, "should be equal")
		})
	}

}
