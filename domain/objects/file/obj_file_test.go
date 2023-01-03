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
	Path:              _fileName,
	GolangPackageName: "content",
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
	require.NoError(t, f.WriteToDisk())
	require.NoError(t, os.Remove(_fileName))
	require.Error(t, f.CheckIfPathExists())
}

func TestWriteToDisk(t *testing.T) {
	require.NoError(t, f.WriteToDisk(), f.CheckIfPathExists())
	require.NoError(t, f.DeletePath())
}

func TestWrite(t *testing.T) {
	content := []byte("main template")
	lenght, errWrite := f.Write(content)
	require.NoError(t, errWrite)
	require.NoError(t, f.CheckIfPathExists())
	require.NoError(t, f.DeletePath())

	assert.Equal(t, lenght, len(content))
}
