package Shared

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func SaveImage(outputImage *image.RGBA64, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, outputImage); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func OpenImage(filename string) (image.Image, error) {
	imageReader, openErr := os.Open(filename)

	if openErr != nil {
		return nil, openErr
	}

	img, _, decodeErr := image.Decode(imageReader)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return img, nil
}

//TODO use alpha as blend percentage
func Blend(c1, c2 color.RGBA64) color.RGBA64 {
	return color.RGBA64{
		R: (c1.R/2 + c2.R/2),
		G: (c1.G/2 + c2.G/2),
		B: (c1.B/2 + c2.B/2),
		A: (c1.A/2 + c2.A/2),
	}
}
