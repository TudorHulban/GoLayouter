package templates

import (
	"html/template"
	"io"
	"os"
)

func renderToPath(templateFilePath, renderToPath string, model any) error {
	file, errCreate := os.Create(renderToPath)
	if errCreate != nil {
		return errCreate
	}
	defer file.Close()

	t, errParse := template.ParseGlob(templateFilePath)
	if errParse != nil {
		return errParse
	}

	return t.Execute(file, model)
}

func RanderTo(templateFilePath string, w io.Writer, model any) error {
	t, errParse := template.ParseGlob(templateFilePath)
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, model)
}
