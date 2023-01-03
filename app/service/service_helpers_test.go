package service

import (
	"log"
	"testing"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/domain/objects/entry"
	"github.com/stretchr/testify/require"
)

func TestCheckIfPathsExists(t *testing.T) {
	happyCases := []struct {
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

	for _, tc := range happyCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRead := helpers.ReadFile(_pathInput + tc.fileInput)
			require.NoError(t, errRead)

			items := entry.NewEntries(content).ParseToItems()

			serv, errNewService := NewService(items)
			require.NoError(t, errNewService)

			require.NoError(t, serv.Render())
			require.NoError(t, serv.CheckIfPathsExists())
			require.NoError(t, serv.DeletePaths())
		})
	}
}

func TestDeletePaths(t *testing.T) {
	happyCases := []struct {
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

	for _, tc := range happyCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRead := helpers.ReadFile(_pathInput + tc.fileInput)
			require.NoError(t, errRead)

			items := entry.NewEntries(content).ParseToItems()

			serv, errNewService := NewService(items)
			require.NoError(t, errNewService)

			require.NoError(t, serv.Render())
			require.NoError(t, serv.CheckIfPathsExists())
			require.NoError(t, serv.DeletePaths())
			require.Error(t, serv.CheckIfPathsExists())
		})
	}
}

func TestChangeDirectory(t *testing.T) {
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
			content, errRead := helpers.ReadFile(_pathInput + tc.fileInput)
			require.NoError(t, errRead)

			entries := entry.NewEntries(content).ParseToItems()

			serv, errNewService := NewService(entries)
			require.NoError(t, errNewService)

			require.NoError(t, helpers.CreateFolder(_temporaryFolder+tc.fileOutput), "creating a folder to write results")
			require.NoError(t, helpers.CheckIfPathExists(_temporaryFolder+tc.fileOutput))

			require.NoError(t, serv.ChangeDirectory(_temporaryFolder+tc.fileOutput))

			require.NoError(t, serv.Render(), "writing error")
			require.NoError(t, serv.CheckIfPathsExists(), "checking error")
			require.NoError(t, serv.DeletePaths())
		})

	}
}

var test = []string{
	"../../temporary_files/program",
	"../../temporary_files/program/helpers1",
	"../../temporary_files/program/helpers1/helpers.go",
	"../../temporary_files/program/helpers1/helpers_file.go",
	"../../temporary_files/program/helpers1/helpers_file_test.go",
	"../../temporary_files/objects1",
	"../../temporary_files/objects1/obj_folder.go",
	"../../temporary_files/objects1/obj_entry.go",
	"../../temporary_files/objects1/obj_entry_test.go",
	"../.	./temporary_files/objects1/obj_file.go",
	"../../temporary_files/objects1/obj_file_test.go",
}

func TestGetPaths(t *testing.T) {
	content, errRead := helpers.ReadFile(_pathInput + "folder_c6")
	require.NoError(t, errRead, "error reading")

	items := entry.NewEntries(content).ParseToItems()
	serv, errNewService := NewService(items)
	require.NoError(t, errNewService)

	paths := serv.GetPaths()
	for _, path := range paths {
		log.Print(path)
	}
	//assert.Equal(t, paths, test)
}
