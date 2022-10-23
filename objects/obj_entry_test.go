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

	require.NoError(t, CreateFilesToDisk(entries), "create files to disk")

	for _, file := range entries {
		assert.Equal(t, helpers.CheckIfFileExists(helpers.RemovePackageName(file)), nil, "check if exists", file)
	}

	for ix := range entries {
		require.NoError(t, RemoveFile(entries[len(entries)-1-ix]))
	}
}
