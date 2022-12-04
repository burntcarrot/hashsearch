package entity

// Image represents the image entity.
type Image struct {
	Path     string
	Distance int
	Hash     uint64
}

// NewImage creates a new image entity.
func NewImage(path string) (*Image, error) {
	// Create image entity.
	img := &Image{
		Path: path,
	}

	// Validate the newly-created image entity.
	err := img.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return img, nil
}

// Validate validates the image entity.
func (img *Image) Validate() error {
	if img.Path == "" {
		return ErrInvalidEntity
	}

	return nil
}
