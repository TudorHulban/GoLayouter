package service

import (
	"errors"
	"fmt"
	"io"

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
		file, err := path.ObjectPath.WriteToDisk()
		if err != nil {
			return fmt.Errorf("error : %w", err)
		}

		if path.Kind != "folder" {
			//TODO : parse an object
			serv.renderFuncs[path.Kind](file, nil)
		}
	}

	return nil
}
