package objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
)

func TestCreateFile(t *testing.T) {
	fileName := "file.go"

	f := &File{
		Path:    fileName,
		Content: "content",
	}

	require.NoError(t, f.WriteToDisk(), helpers.CheckIfFileExists(fileName))

	content, errRead := helpers.ReadFile(fileName)

	require.NoError(t, errRead)
	assert.Equal(t, content[0], f.Content)

	require.NoError(t, RemoveFile(fileName))
}

func TestRemoveFile(t *testing.T) {
	fileName := "file.go"

	f := &File{
		Path:    fileName,
		Content: "",
	}
	require.NoError(t, f.WriteToDisk(), RemoveFile(fileName))
	require.Error(t, helpers.CheckIfFileExists(fileName))
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
			assert.Equal(t, tc.output, helpers.GetFileName(tc.input))
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
		{"invalid path", "../test_cases/folder_c4", "../test_cases/folder_c4_results"},
		{"file without packages", "../test_cases/folder_c5", "../test_cases/folder_c5_results"},
		{"files + paths + packages", "../test_cases/folder_c6", "../test_cases/folder_c6_results"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRe := helpers.ReadFile(tc.fileInput)
			require.NoError(t, errRe, helpers.ClearFile(tc.fileOutput))

			e := NewEntries(content)
			entries := e.Parse()

			require.NoError(t, WriteToFile(entries, tc.fileOutput))

			output, errRead := helpers.ReadFile(tc.fileOutput)
			require.NoError(t, errRead)

			assert.Equal(t, output, entries, "should be equal")
		})
	}

}
