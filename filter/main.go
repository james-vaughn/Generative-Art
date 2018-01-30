package main

import (
	"github.com/fogleman/gg"
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

	context.SavePNG("../output/filter.png")
}

// Draws the background for the image to go over
func drawBackground(context *gg.Context) {
	gradient := gg.NewLinearGradient(0, 0, 1000, 1000)
	gradient.AddColorStop(0, color.White)
	gradient.AddColorStop(1, color.Black)
	context.SetFillStyle(gradient)
	context.DrawRectangle(0, 0, WIDTH, HEIGHT)
	context.Fill()

}
