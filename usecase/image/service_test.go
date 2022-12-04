package image

import (
	"testing"

	"github.com/burntcarrot/hashsearch/entity"
	imageRepository "github.com/burntcarrot/hashsearch/repository/image"
	"github.com/stretchr/testify/assert"
)

func NewFixture() *entity.Image {
	return &entity.Image{
		Path: "testdata/star.png",
	}
}

func TestCreate(t *testing.T) {
	repo := imageRepository.NewInMem()
	service := NewService(repo)
	fixture := NewFixture()

	err := service.Create(fixture.Path)
	assert.Nil(t, err)
}

func TestList(t *testing.T) {
	repo := imageRepository.NewInMem()
	service := NewService(repo)
	fixture := NewFixture()

	_ = service.Create(fixture.Path)
	data, err := service.List()

	assert.Nil(t, err)
	assert.NotZero(t, data)
}

func TestDistances(t *testing.T) {
	repo := imageRepository.NewInMem()
	service := NewService(repo)
	fixture := NewFixture()

	err := service.Create(fixture.Path)
	assert.Nil(t, err)

	err = service.Create("testdata/star-new.png")
	assert.Nil(t, err)

	data, err := service.GetDistances(fixture.Path)

	assert.Nil(t, err)
	assert.NotZero(t, data)

	img1 := data[0]
	img2 := data[1]

	assert.Equal(t, 0, img1.Distance)
	assert.Equal(t, 4, img2.Distance)
}
