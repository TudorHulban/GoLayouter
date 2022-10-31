package objects

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
	"github.com/TudorHulban/GoLayouter/interfaces"
)

const _pathInput = "../test_cases/folder_c6"

func TestConvertToIWritter(t *testing.T) {
	content, errRead := helpers.ReadFile(_pathInput)
	require.NoError(t, errRead)

	e := NewEntries(content)
	entries := e.Parse()

	writter := ConvertToIWritter(entries)

	for _, element := range writter {
		log.Print(element)
	}
}

func TestWrite(t *testing.T) {
	file := File{
		Path:    "directory/main.go",
		Content: "package objects",
	}

	folder := Folder{
		Path: "directory",
	}

	writter := []interfaces.IWritter{folder, file}

	for _, element := range writter {
		element.WriteToDisk()
	}

	require.NoError(t, helpers.CheckIfFileExists(file.Path))
	require.NoError(t, helpers.CheckIfFileExists(folder.Path))

	content, errRead := helpers.ReadFile(file.Path)

	require.NoError(t, errRead, "reading the file path")
	assert.Equal(t, content[0], file.Content)
}
