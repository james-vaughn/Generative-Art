package Shared

import (
	"image"
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
