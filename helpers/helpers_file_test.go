package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const _input = "helpers.in"
const _output = "helpers.out"

func TestTypeofFile(t *testing.T) {
	got := []string{"! .", "# package main", "file.go", "folder"}
	want := []string{"path", "pack", "file", "folder"}

	for i := range got {
		assert.Equal(t, TypeofFile(got[i]), want[i], "verify the type of file")
		fmt.Println(got[i], "is a", want[i])
	}
}

func TestReadFile(t *testing.T) {
	text := []string{"folder1/subfolder1"}

	errClearFile := ClearFile(_input)
	require.NoError(t, errClearFile)

	errWr := WriteTextInFile(text[0], _input)
	require.NoError(t, errWr)

	content, errRe := ReadFile(_input)
	require.NoError(t, errRe)

	assert.Equal(t, text, content, "Text should be equal to read content")
}

func TestClearFile(t *testing.T) {
	text := "this should be deleted"
	
	require.NoError(t, WriteTextInFile(text, _input))

	assert.Equal(t, nil, ClearFile(_input), "should be cleared")
}

func TestReadByLine(t *testing.T) {
	text := []string{"folder1/subfolder1"}
	errClearFile := ClearFile(_input)
	require.NoError(t, errClearFile)

	errWr := WriteTextInFile(text[0], _input)
	require.NoError(t, errWr)

	content, errRe := ReadByLine(_input)
	require.NoError(t, errRe)

	assert.Equal(t, text, content, "Text should be equal to read content")
}

func TestWriteLineInFile(t *testing.T) {
	text := []string{"READ ME ! "}

	errClearFile := ClearFile(_output)
	require.NoError(t, errClearFile)

	errWrite := WriteTextInFile(text[0], _output)
	require.NoError(t, errWrite)

	line, errRl := ReadByLine(_output)
	require.NoError(t, errRl)

	assert.Equal(t, line, text, "Message read from file should be equal to text")
}
