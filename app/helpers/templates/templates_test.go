package templates

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const MainTemplateFilePath = "templates/test"
const randerToPath = "templates/output"

type File struct {
	FileName   string
	Package    string
	IsTestFile bool
}

func TestRenderToPath(t *testing.T) {
	f := File{
		FileName:   "Foo",
		Package:    "test",
		IsTestFile: true,
	}

	require.NoError(t, RanderTo(MainTemplateFilePath, os.Stdout, f))
}
