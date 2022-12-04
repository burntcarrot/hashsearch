package image

import "github.com/burntcarrot/hashsearch/entity"

type Reader interface {
	List() ([]*entity.Image, error)
}

type Writer interface {
	Create(e *entity.Image) error
}

type Repository interface {
	Reader
	Writer
}

type Usecase interface {
	Create(path string) error
	List() ([]*entity.Image, error)
	GetDistances(path string) ([]*entity.Image, error)
}
