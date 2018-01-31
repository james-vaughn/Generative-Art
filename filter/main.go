package main

import (
	"github.com/fogleman/gg"
	noise "github.com/ojrac/opensimplex-go"
	"image/color"
	"os"
	"log"
	"image"
)

const (
	WIDTH  = 1000
	HEIGHT = 1000
)

func main() {

	context := gg.NewContext(WIDTH, HEIGHT)

	drawBackground(context)

	// Overlay image
	imageReader, err := os.Open("../output/termite.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, _ := image.Decode(imageReader)
	context.DrawImage(img, 0, 0)

	context.SavePNG("../output/filter_with_texture.png")
}

// Draws the background for the image to go over
func drawBackground(context *gg.Context) {
	drawSimplexNoiseBackground(context)
}

func drawGradientBackground (context *gg.Context) {
	gradient := gg.NewLinearGradient(0, 0, 1000, 1000)
	gradient.AddColorStop(0, color.White)
	gradient.AddColorStop(1, color.Black)
	context.SetFillStyle(gradient)
	context.DrawRectangle(0, 0, WIDTH, HEIGHT)
	context.Fill()
}

func drawSimplexNoiseBackground (context *gg.Context) {
	noiseGen := noise.NewWithSeed(0) // TODO use random seed
	scale := 10

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			x_val := float64(scale * x)/WIDTH
			y_val := float64(scale * y)/HEIGHT

			n := (noiseGen.Eval2(x_val, y_val) + 1) * (255 / 2)
			noiseVal := uint8(n)
			context.SetColor(color.RGBA{ noiseVal, noiseVal, noiseVal, 255})
			context.SetPixel(x, y)
		}
	}
}