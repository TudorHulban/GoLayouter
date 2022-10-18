package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	var stackFolders Stack

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")

	require.Equal(t, "folder1/subfolder", stackFolders.String())
}

func TestPop(t *testing.T) {
	var stackFolders Stack

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")
	stackFolders.Pop()

	require.Equal(t, "folder1", stackFolders.String())
}

func TestPeek(t *testing.T) {
	var stackFolders Stack

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")

	result := stackFolders.Peek()

	require.Equal(t, "subfolder", result)
}

func TestStringStack(t *testing.T) {
	var stackFolders Stack

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")

	result := stackFolders.String()

	require.Equal(t, "folder1/subfolder", result)
}
