package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeofFile(t *testing.T) {
	got := []string{"! .", "# package main", "file.go", "folder"}
	want := []string{"path", "package", "file", "folder"}

	for i := range got {
		assert.Equal(t, typeofFile(got[i]), want[i], "verify the type of file")
		fmt.Println(got[i], "is a", want[i])
	}
}
