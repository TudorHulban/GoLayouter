package folder

import (
	"os"
	"testing"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const _folderName = "folder"

var f = &Folder{
	Path: _folderName,
}

func TestGetPath(t *testing.T) {
	assert.Equal(t, _folderName, f.Path, "Getting the path")
}

func TestCheckIfPathExist(t *testing.T) {

	_, errWrite := f.WriteToDisk()
	require.NoError(t, errWrite)
	require.NoError(t, f.CheckIfPathExists(), helpers.CheckIfPathExists(_folderName))

	require.NoError(t, f.DeletePath())
	require.Error(t, f.CheckIfPathExists())
}

func TestDeletePath(t *testing.T) {
	_, errWrite := f.WriteToDisk()
	require.NoError(t, errWrite)
	require.NoError(t, errWrite, f.DeletePath())
	require.Error(t, f.CheckIfPathExists())
}

func TestChangeDirectory(t *testing.T) {
	newDirectory := "newDir"
	require.NoError(t, helpers.CreateFolder(newDirectory))

	require.NoError(t, f.ChangeDirectory("newDir"))
	assert.Equal(t, f.Path, newDirectory+"/"+_folderName)

	_, errWrite := f.WriteToDisk()
	require.NoError(t, errWrite)
	require.NoError(t, f.DeletePath())
	require.NoError(t, os.Remove(newDirectory))

	f.Path = (_folderName)
}

func TestWriteToDisk(t *testing.T) {
	_, errWrite := f.WriteToDisk()
	require.NoError(t, errWrite, f.CheckIfPathExists(), f.DeletePath())
}
