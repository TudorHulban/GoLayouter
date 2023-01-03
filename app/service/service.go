package service

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/TudorHulban/GoLayouter/domain/objects/item"
)

type Service struct {
	paths       []item.Item
	renderFuncs map[string]func(io.Writer, any) error
}

func NewService(content []item.Item) (*Service, error) {
	if len(content) == 0 {
		return nil, errors.New("parsed content is empty")
	}

	return &Service{
		paths:       content,
		renderFuncs: _renderFuncs,
	}, nil
}

func (serv Service) Render() error {
	for _, path := range serv.paths {
		errWrite := path.ObjectPath.WriteToDisk()
		if errWrite != nil {
			return fmt.Errorf("error : %w", errWrite)
		}

		if path.Kind != "folder" {
			_, errOpen := os.Open(path.ObjectPath.GetPath())
			if errOpen != nil {
				return errOpen
			}

			//serv.renderFuncs[path.Kind](f, nil)
		}
	}

	return nil
}
