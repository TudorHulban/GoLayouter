package objects

import (
	"log"
	"testing"

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
	content, errRead := helpers.ReadFile(_pathInput)
	require.NoError(t, errRead)

	e := NewEntries(content)
	entries := e.Parse()
	writter := ConvertToIWritter(entries)

	require.NoError(t, interfaces.Write(writter))
	require.NoError(t, interfaces.CheckInterface(writter))
	require.NoError(t, interfaces.DeleteInterface(writter))
}
