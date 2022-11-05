package interfaces

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/objects"
)

func TestWrite(t *testing.T) {
	content, errRead := IRWritterReadFile(_pathInput)
	require.NoError(t, errRead)

	e := objects.NewEntries(content)
	entries := e.Parse()
	writter := objects.ConvertToIWritter(entries)

	require.NoError(t, WriteToDisk(writter))
	require.NoError(t, CheckPathsExists(writter))
	require.NoError(t, DeletePaths(writter))
}

func TestConvertToIWritter(t *testing.T) {
	content, errRead := IRWritterReadFile(_pathInput)
	require.NoError(t, errRead)

	e := objects.NewEntries(content)
	entries := e.Parse()
	writter := objects.ConvertToIWritter(entries)

	for _, element := range writter {
		log.Print(element)
	}
}
