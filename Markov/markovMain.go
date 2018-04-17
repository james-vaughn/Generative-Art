package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
)

func main() {
	image, err := openImage("")

	if err != nil {
		log.Fatalf("Could not open image: %v", err)
	}

	markovChain := makeChain(image)

	fmt.Println(markovChain)
	_ = image
}

func openImage(filename string) (image.Image, error) {
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

func makeChain(image image.Image) map[color.RGBA]map[color.RGBA]int {
	markovChain := make(map[color.RGBA]map[color.RGBA]int)

	for col := image.Bounds().Min.Y; col < image.Bounds().Max.Y; col++ {
		for row := image.Bounds().Min.X; row < image.Bounds().Max.X; row++ {
			log.Println(image.At(row, col))
			//inner := make(map[color.RGBA]int)
			//inner[color.RGBA{2, 1, 1, 1}] = 1
			//markovChain[color.RGBA{1, 1, 1, 1}] = inner
		}
	}

	return markovChain
}
