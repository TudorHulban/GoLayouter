package templates

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
)

const from = "input/"
const to = "test/"

type Model struct {
	FileName   string
	ObjectName string
	Package    string
}

var m = Model{
	FileName:   "main.go",
	ObjectName: "Main",
	Package:    "main",
}

func TestRenderToPath(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		output      string
	}{
		{"main template", from + "main", to + "main_output"},
		{"object template", from + "object", to + "object_output"},
		{"tdd template", from + "tableDriven", to + "tdd_output"},
		{"test template", from + "test", to + "test_output"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			{
				buffer, errCreate := os.Create("buffer")
				require.NoError(t, errCreate)
				require.NoError(t, RenderTo(buffer, tc.input, m))

				cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode
				equal, errCompare := cmp.CompareFile(buffer.Name(), tc.output)

				require.NoError(t, errCompare)
				require.Equal(t, true, equal)
				require.NoError(t, os.Remove(buffer.Name()))
			}
		})
	}
}
