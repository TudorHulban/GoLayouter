package interfaces

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/domain/objects"
)

func TestWrite(t *testing.T) {
	content, errRead := IRWritterReadFile(_pathInput)
	require.NoError(t, errRead)

	e := objects.NewEntries(content)
	entries := e.Parse()
	writter := ConvertToIWritter(entries)

	require.NoError(t, WriteToDisk(writter), "writing error")
	require.NoError(t, CheckPathsExists(writter), "checking error")
	require.NoError(t, DeletePaths(writter), "deleting error")
}

func TestConvertToIWritter(t *testing.T) {
	content, errRead := IRWritterReadFile(_pathInput)
	require.NoError(t, errRead, "error reading")

	e := objects.NewEntries(content)
	entries := e.Parse()

	writter := ConvertToIWritter(entries)

	for _, element := range writter {
		log.Print(element)
	}
}
