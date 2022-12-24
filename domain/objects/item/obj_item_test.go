package item

import (
	"log"
	"testing"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	"github.com/TudorHulban/GoLayouter/domain/objects/entry"
	"github.com/stretchr/testify/require"
)

const _TestCases = "../../test_cases/"

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
			entries := entry.NewEntries(content).Parse()

			for _, item := range entries {
				log.Print(item.path, " ", item.kind)
			}

		})
	}
}
