// TODO: rename this file
package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/app/helpers"
)

const _fileName = "file.go"

var f = &File{
	Path:    _fileName,
	Content: "content",
}

func TestCheckIfPathExists(t *testing.T) {
	require.NoError(t, f.WriteToDisk())
	require.NoError(t, f.CheckIfPathExists(), helpers.CheckIfPathExists(_fileName))

	require.NoError(t, f.DeletePath())
	require.Error(t, f.CheckIfPathExists())
}

func TestChangeDirectory(t *testing.T) {
	newDirectory := "newDir"
	require.NoError(t, helpers.CreateFolder(newDirectory))

	require.NoError(t, f.ChangeDirectory("newDir"))
	assert.Equal(t, f.Path, newDirectory+"/"+_fileName)

	require.NoError(t, f.WriteToDisk())
	require.NoError(t, f.DeletePath())
	require.NoError(t, os.Remove(newDirectory))

	f.Path = (_fileName)
}

func TestDeletePath(t *testing.T) {
	require.NoError(t, f.WriteToDisk(), os.Remove(_fileName))
	require.Error(t, f.CheckIfPathExists())
}

func TestWriteToDisk(t *testing.T) {
	require.NoError(t, f.WriteToDisk(), f.CheckIfPathExists())

	content, errRead := helpers.ReadFile(_fileName)

	require.NoError(t, errRead)
	assert.Equal(t, content[0], f.Content)

	require.NoError(t, f.DeletePath())
}
