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
