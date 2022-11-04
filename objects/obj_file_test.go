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

	require.NoError(t, f.WriteToDisk(), helpers.CheckIfPathExists(fileName))

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
	require.Error(t, helpers.CheckIfPathExists(fileName))
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
