package service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

const _pathInput = "../../test_cases/folder_c6"

func TestConvertToIWritter(t *testing.T) {
	content, errRead := helpers.ReadFile(_pathInput)
	require.NoError(t, errRead, "error reading")

	e := objects.NewEntries(content)
	entries := e.Parse()

	writter := service.ConvertToIWritter(entries)

	for _, element := range writter {
		log.Print(element)
	}
}

func TestWrite(t *testing.T) {
	content, errRead := helpers.ReadFile(_pathInput)
	require.NoError(t, errRead)

	e := objects.NewEntries(content)
	entries := e.Parse()
	writter := service.ConvertToIWritter(entries)

	require.NoError(t, service.WriteToDisk(writter), "writing error")
	require.NoError(t, CheckPathsExists(writter), "checking error")
	require.NoError(t, DeletePaths(writter), "deleting error")
}
