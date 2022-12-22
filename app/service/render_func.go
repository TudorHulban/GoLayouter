package service

import (
	"io"

	"github.com/TudorHulban/GoLayouter/app/helpers/templates"
)

const (
	_templateMain        = "../templates/templates/main"
	_templateTest        = "../templates/templates/test"
	_templateObject      = "../templates/templates/object"
	_templateTableDriven = "../templates/temlates/tableDriven"
)

var _renderFuncs = map[string]func(io.Writer, any) error{
	"main":        renderMain,
	"test":        renderTest,
	"object":      renderObject,
	"tableDriven": renderTableDriven,
}

func renderMain(w io.Writer, _ any) error {
	return templates.RanderTo(_templateMain, w, nil)
}

func renderTest(w io.Writer, object any) error {
	return templates.RanderTo(_templateTest, w, object)
}

func renderObject(w io.Writer, object any) error {
	return templates.RanderTo(_templateObject, w, object)
}

func renderTableDriven(w io.Writer, object any) error {
	return templates.RanderTo(_templateTableDriven, w, object)
}
