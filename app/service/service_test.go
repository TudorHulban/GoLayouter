package service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

const _pathInput = "../../test_cases/folder_c6"

func TestConvertToIFileOperations(t *testing.T) {
	content, errRead := helpers.ReadFile(_pathInput)
	require.NoError(t, errRead, "error reading")

	entries := objects.NewEntries(content).Parse()

	var serv Service

	writter := serv.ConvertToIFileOperations(entries)

	for _, element := range writter {
		log.Print(element)
	}
}

func TestWrite(t *testing.T) {
	content, errRead := helpers.ReadFile(_pathInput)
	require.NoError(t, errRead)

	entries := objects.NewEntries(content).Parse()

	serv := NewService(entries)

	require.NoError(t, serv.WriteToDisk(), "writing error")
	require.NoError(t, serv.CheckPathsExists(), "checking error")
	require.NoError(t, serv.DeletePaths(), "deleting error")
}
