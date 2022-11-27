package objects

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
)

const _fileName = "file.go"

func TestCreateFile(t *testing.T) {
	f := &File{
		Path:    _fileName,
		Content: "content",
	}

	require.NoError(t, f.WriteToDisk(), helpers.CheckIfPathExists(_fileName))

	content, errRead := helpers.ReadFile(_fileName)

	require.NoError(t, errRead)
	assert.Equal(t, content[0], f.Content)

	require.NoError(t, RemoveFile(_fileName))
}

func TestRemoveFile(t *testing.T) {
	f := &File{
		Path:    _fileName,
		Content: "",
	}
	require.NoError(t, f.WriteToDisk(), RemoveFile(_fileName))
	require.Error(t, helpers.CheckIfPathExists(_fileName))
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
			_, fileName := path.Split(tc.input)
			assert.Equal(t, tc.output, fileName)
		})
	}
}
