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

func (serv Service) WriteToDisk() error {
	for _, path := range serv.paths {
		if err := path.ObjectPath.WriteToDisk(); err != nil {
			return fmt.Errorf("error : %w", err)
		}

		// TODO renter by path.Kind
	}

	return nil
}

func (serv *Service) ChangeDirectory(newPath string) error {
	for _, path := range serv.paths {
		if err := path.ObjectPath.ChangeDirectory(newPath); err != nil {
			return fmt.Errorf("error : %w", err)
		}
	}

	return nil
}
