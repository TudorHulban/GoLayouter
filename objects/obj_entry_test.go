package objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
)

const _pathInput = "../test_cases/folder_c1"

func TestWriteToDisk(t *testing.T) {
	content, errRe := helpers.ReadFile(_pathInput)
	require.NoError(t, errRe)

	e := NewEntries(content)

	entries := e.Parse()

	errCreate := CreateFilesToDisk(entries)
	require.NoError(t, errCreate)

	for _, file := range entries {
		err := helpers.CheckIfExist(file)
		assert.Equal(t, err, nil, "No match for file", file)
	}

	for i := range entries {
		errRemove := RemoveFile(entries[len(entries)-1-i])
		require.NoError(t, errRemove)
	}
}
