package objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
)

func TestWriteToDisk(t *testing.T) {
	content, errRe := helpers.ReadFile(helpers.PathInput)
	require.NoError(t, errRe)

	e := NewEntries(content)

	entries := e.Parse()
	CreateFilesToDisk(entries)

	for _, file := range entries {
		err := helpers.CheckIfExist(file)
		assert.Equal(t, err, nil, "No match for file", file)
	}

	for i, _ := range entries {
		errRemove := RemoveFile(entries[len(entries)-1-i])
		require.NoError(t, errRemove)
	}
}
