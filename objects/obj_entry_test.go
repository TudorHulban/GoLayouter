package objects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TudorHulban/GoLayouter/helpers"
)

func WriteToFile(entries []string, output string) error {
	for _, file := range entries {
		if err := helpers.WriteTextInFile(file, output); err != nil {
			return err
		}
	}

	return nil
}

func TestWriteObjectsToFile(t *testing.T) {
	testCases := []struct {
		description string
		fileInput   string
		fileOutput  string
	}{
		{"2 levels", "../test_cases/folder_c1", "../test_cases/folder_c1_results"},
		{"3 levels", "../test_cases/folder_c2", "../test_cases/folder_c2_results"},
		{"3 levels with going back", "../test_cases/folder_c3", "../test_cases/folder_c3_results"},
		{"invalid path", "../test_cases/folder_c4", "../test_cases/folder_c4_results"},
		{"file without packages", "../test_cases/folder_c5", "../test_cases/folder_c5_results"},
		{"files + paths + packages", "../test_cases/folder_c6", "../test_cases/folder_c6_results"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRe := helpers.ReadFile(tc.fileInput)
			require.NoError(t, errRe, helpers.ClearFile(tc.fileOutput))

			e := NewEntries(content)
			entries := e.Parse()

			require.NoError(t, WriteToFile(entries, tc.fileOutput))

			output, errRead := helpers.ReadFile(tc.fileOutput)
			require.NoError(t, errRead)

			assert.Equal(t, output, entries, "should be equal")
		})
	}
}
