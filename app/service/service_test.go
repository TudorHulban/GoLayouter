package service

import (
	"testing"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/domain/objects/entry"
	cases "github.com/TudorHulban/GoLayouter/test_cases"
	"github.com/stretchr/testify/require"
)

const _pathInput = "../../test_cases/files/"
const _temporaryFolder = "../../temporary_files/"

func TestWrite(t *testing.T) {
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
		})

	}

	for _, tc := range cases.TestCasesInvalid {
		t.Run(tc.Description, func(t *testing.T) {
			content, errRead := helpers.ReadFile(_pathInput + tc.FileInput)
			require.NoError(t, errRead)

			items := entry.NewEntries(content).ParseToItems()

			serv, errNewService := NewService(items)
			require.NoError(t, errNewService)

			require.NoError(t, helpers.CreateFolder(_temporaryFolder+tc.FileOutput), "creating a folder to write results")
			require.NoError(t, helpers.CheckIfPathExists(_temporaryFolder+tc.FileOutput))

			require.NoError(t, serv.ChangeDirectory(_temporaryFolder+tc.FileOutput))

			require.Error(t, serv.Render(), "writing error")
			require.Error(t, serv.CheckIfPathsExists(), "checking error")
		})

	}
}
