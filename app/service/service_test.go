package service

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/app/helpers/helpers"
	"github.com/TudorHulban/GoLayouter/domain/objects"
)

const _pathInput = "../test_cases/folder_c6"
const _temporaryFolder = "../../temporary_files/"

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
	testCases := []struct {
		description string
		fileInput   string
		fileOutput  string
	}{
		{"2 levels", "folder_c1", "folder_c1_results"},
		{"3 levels", "folder_c2", "folder_c2_results"},
		{"3 levels with going back", "folder_c3", "folder_c3_results"},

		{"file without packages", "folder_c5", "folder_c5_results"},
		{"files + paths + packages", "folder_c6", "folder_c6_results"},
		{"small test with packages", "folder_c7", "folder_c7_results"},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRead := helpers.ReadFile("../../test_cases/" + tc.fileInput)
			require.NoError(t, errRead)

			entries := objects.NewEntries(content).Parse()

			serv, errNewService := NewService(entries)
			require.NoError(t, errNewService)

			require.NoError(t, helpers.CreateFolder(_temporaryFolder+tc.fileOutput), "creating a folder to write results")
			require.NoError(t, helpers.CheckIfPathExists(_temporaryFolder+tc.fileOutput))

			require.NoError(t, serv.ChangeDirectory(_temporaryFolder+tc.fileOutput))

			require.NoError(t, serv.WriteToDisk(), "writing error")
			require.NoError(t, serv.CheckPathsExists(), "checking error")
		})

	}
	testCases = []struct {
		description string
		fileInput   string
		fileOutput  string
	}{
		{"invalid path", "folder_c4", "folder_c4_results"},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRead := helpers.ReadFile("../../test_cases/" + tc.fileInput)
			require.NoError(t, errRead)

			entries := objects.NewEntries(content).Parse()

			serv, errNewService := NewService(entries)
			require.NoError(t, errNewService)

			require.NoError(t, helpers.CreateFolder(_temporaryFolder+tc.fileOutput), "creating a folder to write results")
			require.NoError(t, helpers.CheckIfPathExists(_temporaryFolder+tc.fileOutput))

			require.NoError(t, serv.ChangeDirectory(_temporaryFolder+tc.fileOutput))

			require.Error(t, serv.WriteToDisk(), "writing error")
			require.Error(t, serv.CheckPathsExists(), "checking error")
		})

	}
}
