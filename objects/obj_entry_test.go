package objects

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
)

const _pathInput = "../test_cases/folder_c6"

func TestContentSlicing(t *testing.T) {
	content, errRead := helpers.ReadFile(_pathInput)
	require.NoError(t, errRead)

	e := NewEntries(content)
	entries := e.Parse()

	files, folder := ContentSlicing(entries)

	for _, element := range files {
		log.Print(element)
	}

	for _, element := range folder {
		log.Print(element)
	}
}
