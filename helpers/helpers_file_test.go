package helpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	content, errRe := ReadFile(_filePath)
	require.NoError(t, errRe)

	// TODO: add test

	fmt.Println(strings.Join(content, "\n"))
}

// TODO: add struct name "add"
