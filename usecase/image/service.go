package image

import (
	"github.com/burntcarrot/hashsearch/entity"
	"github.com/burntcarrot/hashsearch/pkg/hash"
)

type Service struct {
	repo Repository
}

// NewService returns the image service.
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(path string) error {
	// Create image entity.
	e, err := entity.NewImage(path)
	if err != nil {
		return err
	}

	// Load as image.
	img, err := hash.LoadAsImage(e.Path)
	if err != nil {
		return err
	}

	// Set the entity's hash value using the loaded image.
	ah := hash.NewAverageHash()
	e.Hash = ah.GenerateHash(img)

	return s.repo.Create(e)
}

func (s *Service) List() ([]*entity.Image, error) {
	// Get all images.
	images, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, entity.ErrNotFound
	}

	return images, nil
}

func (s *Service) GetDistances(path string) ([]*entity.Image, error) {
	// Load source image.
	img, err := hash.LoadAsImage(path)
	if err != nil {
		return nil, err
	}

	// Get the source image's hash.
	ah := hash.NewAverageHash()
	sourceImageHash := ah.GenerateHash(img)

	// Get all images.
	images, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(images) == 0 {
		return nil, entity.ErrNotFound
	}

	// Calculate distances for each image with respect to the source image.
	for _, img := range images {
		img.Distance = hash.CalculateDistance(sourceImageHash, img.Hash)
	}

	return images, nil
}
