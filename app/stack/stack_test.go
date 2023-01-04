package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack_IsEmpty(t *testing.T) {
	var stackFolders Stack

	require.Equal(t, true, stackFolders.IsEmpty(), "is empty testing")

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")

	require.NotEqual(t, true, stackFolders.IsEmpty(), "is empty testing")

}

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
	stackFolders.Push("subsubsub")

	stackFolders.Pop()

	require.Equal(t, "folder1/subfolder", stackFolders.String())
}

func TestPeek(t *testing.T) {
	var stackFolders Stack

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")
	stackFolders.Push("subsubsub")

	result := stackFolders.Peek()

	require.Equal(t, "subsubsub", result)
	require.Equal(t, stackFolders.String(), "folder1/subfolder/subsubsub")
}

func TestStringStack(t *testing.T) {
	var stackFolders Stack

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")

	result := stackFolders.String()

	require.Equal(t, "folder1/subfolder", result)
}
