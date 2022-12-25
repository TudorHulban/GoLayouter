package entry

import (
	"log"
	"testing"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const _TestCases = "../../../test_cases/"

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
		{"2 levels", "folder_c1", "folder_c1_results"},
		{"3 levels", "folder_c2", "folder_c2_results"},
		{"3 levels with going back", "folder_c3", "folder_c3_results"},
		{"invalid path", "folder_c4", "folder_c4_results"},
		{"file without packages", "folder_c5", "folder_c5_results"},
		{"files + paths + packages", "folder_c6", "folder_c6_results"},
		{"small test with packages", "folder_c7", "folder_c7_results"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRe := helpers.ReadFile(_TestCases + tc.fileInput)
			require.NoError(t, errRe)
			entries := NewEntries(content).ParseToStrings()

			output, errRead := helpers.ReadFile(_TestCases + tc.fileOutput)
			require.NoError(t, errRead)

			assert.Equal(t, output, entries, "should be equal")
		})
	}
}

func TestParseToItems(t *testing.T) {
	testCases := []struct {
		description string
		fileInput   string
		fileOutput  string
	}{
		{"files + paths + packages", "folder_c6", "folder_c6_results"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			content, errRe := helpers.ReadFile(_TestCases + tc.fileInput)
			require.NoError(t, errRe)
			entries := NewEntries(content).ParseToItems()

			for _, item := range entries {
				log.Print(item.ObjectPath, item.Kind)
			}

		})
	}
}
