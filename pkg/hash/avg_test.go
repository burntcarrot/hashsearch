package hash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage_GenerateHash(t *testing.T) {
	tests := []struct {
		description string
		path        string
		hash        string
	}{
		{
			description: "normal image (example 1)",
			path:        "../../usecase/image/testdata/star-new.png",
			hash:        "0001000000110000111100001111110011111100111100000011000000010000",
		},
		{
			description: "normal image (example 2)",
			path:        "../../usecase/image/testdata/star.png",
			hash:        "0000000000010000111100001111110011111100111100000001000000000000",
		},
	}

	ah := NewAverageHash()

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			img, err := LoadAsImage(tc.path)
			if err != nil {
				t.Error(err)
			}

			hash := fmt.Sprintf("%064b", ah.GenerateHash(img))
			assert.Equal(t, tc.hash, hash)
		})
	}
}

func TestAverage_CalculateDistance(t *testing.T) {
	tests := []struct {
		description string
		img1        string
		img2        string
		distance    int
	}{
		{
			description: "same images",
			img1:        "../../usecase/image/testdata/star.png",
			img2:        "../../usecase/image/testdata/star.png",
			distance:    0,
		},
		{
			description: "original vs edited image",
			img1:        "../../usecase/image/testdata/star.png",
			img2:        "../../usecase/image/testdata/star-new.png",
			distance:    4,
		},
	}

	ah := NewAverageHash()

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			img1, err := LoadAsImage(tc.img1)
			if err != nil {
				t.Error(err)
			}

			hash1 := ah.GenerateHash(img1)

			img2, err := LoadAsImage(tc.img2)
			if err != nil {
				t.Error(err)
			}

			hash2 := ah.GenerateHash(img2)

			distance := CalculateDistance(hash1, hash2)
			assert.Equal(t, tc.distance, distance)
		})
	}
}
