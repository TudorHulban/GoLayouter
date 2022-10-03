package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {
	content, errRe := readFile(_filePath)
	require.NoError(t, errRe)

	entries := parse(convertToEntries(content))

	for _, file := range entries {
		err := checkIfExist(file)
		assert.Equal(t, err, nil, "No match for file", file)
	}

}

func TestTypeofFile(t *testing.T) {
	got := []string{"! .", "# package main", "file.go", "folder"}
	want := []string{"path", "package", "file", "folder"}

	for i := range got {
		assert.Equal(t, typeofFile(got[i]), want[i], "verify the type of file")
		fmt.Println(got[i], "is a", want[i])
	}
}
