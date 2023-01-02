package item

import (
	"github.com/TudorHulban/GoLayouter/domain"
)

type Item struct {
	ObjectPath domain.IFileOperations

	Kind string
}
