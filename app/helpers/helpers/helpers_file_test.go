package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const _input = "helpers.in"
const _output = "helpers.out"

func TestRemovePackageName(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		output      string
	}{
		{"with package", "program/objects/obj_folder.go(objectes)", "program/objects/obj_folder.go"},
		{"without package & only folders", "folder-root1/subfolder1", "folder-root1/subfolder1"},
		{"without package + files", "folder-root1/subfolder1/main.go", "folder-root1/subfolder1/main.go"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.output, RemovePackageName(tc.input))
		})
	}
}

func TestParsePackage(t *testing.T) {
	got := "program/objects/obj_folder.go(objectes)"
	want := "objectes"

	assert.Equal(t, ParsePackage(got), want, "test parse")
}

func TestTypeofFile(t *testing.T) {
	got := []string{"! .", "# package main", "file.go", "folder"}
	want := []string{"path", "pack", "file", "folder"}

	for i := range got {
		assert.Equal(t, TypeofFile(got[i]), want[i], "verify the type of file")
		//fmt.Println(got[i], "is a", want[i])
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

	assert.Equal(t, nil, ClearFile(_input), "should be empty")
}

func TestWriteTextInFile(t *testing.T) {
	text := []string{"READ ME ! "}

	errClearFile := ClearFile(_output)
	require.NoError(t, errClearFile)

	errWrite := WriteTextInFile(text[0], _output)
	require.NoError(t, errWrite)

	line, errRl := ReadFile(_output)
	require.NoError(t, errRl)

	assert.Equal(t, line, text, "Message read from file should be equal to text")
}

func TestFileExists(t *testing.T) {
	err1 := CheckIfPathExists("x")
	require.Error(t, err1, err1) // stat x: no such file or directory

	err2 := CheckIfPathExists("/var")
	require.NoError(t, err2, err2)
}

func TestGetCommand(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		output      string
	}{
		{"package", "# package", "package"},
		{"path", "! home/path", "home/path"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.output, GetCommand(tc.input))
		})
	}
}

func TestIsTestFile(t *testing.T) {
	assert.Equal(t, true, IsTestFile("t"))
	assert.NotEqual(t, true, IsTestFile("something"))
}

func TestCreateGolangTestFile(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		checkResult func(string, error)
	}{
		{"1 happy path", "main.go", func(res string, err error) {
			assert.NoError(t, err)
			assert.Equal(t, "main_test.go", res)
		}},
		{"2", "functions.go", func(res string, err error) {
			assert.NoError(t, err)
			assert.Equal(t, "functions_test.go", res)
		}},
		{"3 typo", "../main..go", func(res string, err error) {
			assert.NoError(t, err)
			assert.Equal(t, "../main_test.go", res)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			tc.checkResult(CreateGolangTestFile(tc.input))
		})
	}
}

func TestCheckIfPathExists(t *testing.T) {
	require.NoError(t, CheckIfPathExists(_input))
	require.NoError(t, CheckIfPathExists(_output))
	require.Error(t, CheckIfPathExists("invalid-path"))
}

func TestConvertToFiles(t *testing.T) {
	input := "file.go main.go head.go"
	packageName := "t"

	want := []string{"file.go", "file_test.go", "main.go", "main_test.go", "head.go", "head_test.go"}

	assert.Equal(t, ConvertToFiles(input, packageName), want)
}
