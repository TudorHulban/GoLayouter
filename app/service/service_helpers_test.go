package service

import (
	"testing"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/domain/objects/entry"
	cases "github.com/TudorHulban/GoLayouter/test_cases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckIfPathsExists(t *testing.T) {
	for _, tc := range cases.TestCasesHappyPath {
		t.Run(tc.Description, func(t *testing.T) {
			content, errRead := helpers.ReadFile(_pathInput + tc.FileInput)
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
	for _, tc := range cases.TestCasesHappyPath {

		t.Run(tc.Description, func(t *testing.T) {
			content, errRead := helpers.ReadFile(_pathInput + tc.FileInput)
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
	for _, tc := range cases.TestCasesHappyPath {
		t.Run(tc.Description, func(t *testing.T) {
			content, errRead := helpers.ReadFile(_pathInput + tc.FileInput)
			require.NoError(t, errRead)

			entries := entry.NewEntries(content).ParseToItems()

			serv, errNewService := NewService(entries)
			require.NoError(t, errNewService)

			require.NoError(t, helpers.CreateFolder(_temporaryFolder+tc.FileOutput), "creating a folder to write results")
			require.NoError(t, helpers.CheckIfPathExists(_temporaryFolder+tc.FileOutput))

			require.NoError(t, serv.ChangeDirectory(_temporaryFolder+tc.FileOutput))

			require.NoError(t, serv.Render(), "writing error")
			require.NoError(t, serv.CheckIfPathsExists(), "checking error")
			require.NoError(t, serv.DeletePaths())
		})

	}
}

func TestGetPaths(t *testing.T) {
	for _, tc := range cases.TestCasesHappyPath {
		t.Run(tc.Description, func(t *testing.T) {
			content, errRead := helpers.ReadFile(_pathInput + tc.FileInput)
			require.NoError(t, errRead, "error reading")

			items := entry.NewEntries(content).ParseToItems()
			serv, errNewService := NewService(items)
			require.NoError(t, errNewService)

			test, errR := helpers.ReadFile(_pathInput + tc.FileOutput)
			require.NoError(t, errR, "error reading")

			paths := serv.GetPaths()
			assert.Equal(t, paths, test)
		})
	}
}
