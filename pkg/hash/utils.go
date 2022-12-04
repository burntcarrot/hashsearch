package hash

import (
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

// Resize resizes the image and returns it.
func Resize(width, height uint, img image.Image) image.Image {
	dst := resize.Resize(width, height, img, resize.NearestNeighbor)
	return dst
}

// Grayscale converts image to grayscale.
func Grayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rgba := img.At(x, y)
			grayImg.Set(x, y, rgba)
		}
	}

	return grayImg
}

// FindMean finds the mean from the image.
func FindMean(img image.Image) uint32 {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	var sum uint32
	totalPixels := uint32(width * height)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			sum += r
		}
	}

	return sum / totalPixels
}

// LoadAsImage loads image from the path.
func LoadAsImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return img, nil
}

// CalculateDistance calculates the hamming distance from two hashes.
func CalculateDistance(x, y uint64) int {
	distance := 0
	x ^= y
	for x > 0 {
		distance += 1
		x &= x - 1
	}

	return distance
}
