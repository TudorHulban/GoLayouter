// TODO: rename this file
package objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	helpers "github.com/TudorHulban/GoLayouter/app/helpers/utils"
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
