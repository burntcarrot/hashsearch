package repository

import "github.com/burntcarrot/hashsearch/entity"

// InMem is the repository using the memory.
type InMem struct {
	m map[string]*entity.Image
}

// NewInMem creates a new repository using the memory.
func NewInMem() *InMem {
	var m = map[string]*entity.Image{}
	return &InMem{
		m: m,
	}
}

// Create creates an image entry.
func (r *InMem) Create(e *entity.Image) error {
	r.m[e.Path] = e
	return nil
}

// List returns a list of image entries.
func (r *InMem) List() ([]*entity.Image, error) {
	var data []*entity.Image

	for _, d := range r.m {
		data = append(data, d)
	}

	return data, nil
}
