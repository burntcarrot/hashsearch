package hash

import (
	"image"
)

type AvgHash struct{}

func NewAverageHash() *AvgHash {
	return &AvgHash{}
}

// GenerateHash generates the hash value using average hash.
func (ah *AvgHash) GenerateHash(img image.Image) uint64 {
	// Resize source image to 8x8.
	resizedImg := Resize(8, 8, img)

	// Convert the resized image to grayscale.
	grayImg := Grayscale(resizedImg)

	// Find the mean using the grayscale image.
	mean := FindMean(grayImg)

	// Return average hash.
	return ah.getHash(grayImg, mean)
}

// getHash generates the average hash from the grayscale image and the mean.
func (ah *AvgHash) getHash(img image.Image, mean uint32) uint64 {
	var hash, pos uint64
	pos = 1
	bounds := img.Bounds()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			if r > mean {
				hash |= 1 << pos // set bit to 1
			}
			pos++
		}
	}

	return hash
}
