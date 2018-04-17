package main

import (
	"image"
	_ "image/png" //needed to decode png
	"log"
	"math"
	"os"

	gg "github.com/fogleman/gg"
)

const (
	WIDTH  = 1080
	HEIGHT = 1080
)

//https://jonnoftw.github.io/2017/01/18/markov-chain-image-generation
func main() {
	image, err := openImage("input/test.png")

	if err != nil {
		log.Fatalf("Could not open image: %v", err)
	}

	markovChain := makeChain(image)
	//fmt.Println(markovChain)

	context := gg.NewContext(WIDTH, HEIGHT)
	drawImage(context, markovChain)

	context.SavePNG("output/markov.png")
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

func drawImage(context *gg.Context, chain *MarkovChain) {
	var x, y int

	for y = 0; y < WIDTH; y++ {
		for x = 0; x < HEIGHT; x++ {
			c := chain.Next()

			//fmt.Println(c)

			context.SetColor(toRGBA(c))

			context.SetPixel(x, y)
		}
	}
}

func makeChain(image image.Image) *MarkovChain {
	markovChain := NewChain()

	minY := image.Bounds().Min.Y
	minX := image.Bounds().Min.X
	maxY := image.Bounds().Max.Y
	maxX := image.Bounds().Max.X

	for col := minY; col < maxY; col++ {
		for row := minX; row < maxX; row++ {
			markovChain.Add(
				image.At(clamp(row, minX, maxX), clamp(col, minY, maxY)),
				image.At(clamp(row, minX, maxX), clamp(col+1, minY, maxY)))

			markovChain.Add(
				image.At(clamp(row, minX, maxX), clamp(col, minY, maxY)),
				image.At(clamp(row, minX, maxX), clamp(col-1, minY, maxY)))

			markovChain.Add(
				image.At(clamp(row, minX, maxX), clamp(col, minY, maxY)),
				image.At(clamp(row+1, minX, maxX), clamp(col, minY, maxY)))

			markovChain.Add(
				image.At(clamp(row, minX, maxX), clamp(col, minY, maxY)),
				image.At(clamp(row-1, minX, maxX), clamp(col, minY, maxY)))
		}
	}

	return markovChain
}

func clamp(x, min, max int) int {
	return int(math.Max(float64(min), math.Min(float64(max), float64(x))))
}
