package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/udhos/equalfile"
)

const _toTest = "../templates/test/"

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

func TestRenderFuncs(t *testing.T) {
	testCases := []struct {
		description string
		testOutput  string
		kind        string
		object      Model
	}{
		{"Render Main", "main", "main", m},
		{"Render Object", "object", "object", m},
		{"Render TDD", "tdd", "tableDriven", m},
		{"Render Test", "test", "test", m},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			{
				buffer, errCreate := os.Create("buffer")
				require.NoError(t, errCreate)

				_renderFuncs[tc.kind](buffer, tc.object)
				cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode
				equal, errCompare := cmp.CompareFile(buffer.Name(), _toTest+tc.testOutput+"_output")
				require.NoError(t, errCompare)

				require.Equal(t, true, equal)
				require.NoError(t, os.Remove(buffer.Name()))
			}
		})
	}
}
