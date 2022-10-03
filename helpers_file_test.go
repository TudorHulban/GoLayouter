package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	content, errRe := readFile(_filePath)
	require.NoError(t, errRe)

	// TODO: add test

	fmt.Println(strings.Join(content, "\n"))
}
