package entry

import (
	"testing"

	"github.com/TudorHulban/GoLayouter/app/helpers"
	cases "github.com/TudorHulban/GoLayouter/test_cases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const _TestCases = "../../../test_cases/files/"

func TestParseToItems(t *testing.T) {
	for _, tc := range cases.TestCasesHappyPath {
		t.Run(tc.Description, func(t *testing.T) {
			content, errRe := helpers.ReadFile(_TestCases + tc.FileInput)
			require.NoError(t, errRe)
			items := NewEntries(content).ParseToItems()

			var test []string
			for _, item := range items {
				test = append(test, item.ObjectPath.GetPath())
			}

			output, errR := helpers.ReadFile(_TestCases + tc.FileOutput)
			require.NoError(t, errR, "error reading")

			assert.Equal(t, test, output, "should be equal")
		})
	}
}
