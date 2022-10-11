package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	var stackFolders Stack

	stackFolders.Push("folder1")
	stackFolders.Push("subfolder")

	require.Equal(t, "folder1/subfolder", stackFolders.String())
}
