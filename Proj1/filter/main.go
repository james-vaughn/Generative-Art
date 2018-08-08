package main

import (
	"github.com/fogleman/gg"
	noise "github.com/ojrac/opensimplex-go"
	"image/color"
	"os"
	"log"
	"image"
	"time"
)

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

func main() {
	log.Println("Applying filter...")

	context := gg.NewContext(WIDTH, HEIGHT)

	drawBackground(context)
	context.SavePNG("output/filter.png")

	// Overlay images
	err := overlayImages(context, []string {
		//"output/termite.png",
		//"output/termite2.png",
	})

	if err != nil {
		log.Fatal(err)
	}

	context.SavePNG("output/filter_with_texture.png")
}

func overlayImages(context *gg.Context, imageNames []string)  error {
	for _, imageFileName := range imageNames {
		if err := overlayImage(context, imageFileName); err != nil {
			return err
		}
	}

	return nil
}

func overlayImage(context *gg.Context, filename string) error {
	imageReader, openErr := os.Open(filename)

	if openErr != nil {
		return openErr
	}

	img, _, decodeErr := image.Decode(imageReader)

	if decodeErr != nil {
		return decodeErr
	}

	context.DrawImage(img, 0, 0)
	
	return nil
}

// Draws the background for the image to go over
func drawBackground(context *gg.Context) {
	// drawGradientBackground(context)	
	drawSimplexNoiseBackground(context)
}

func drawGradientBackground (context *gg.Context) {
	gradient := gg.NewLinearGradient(0, 0, WIDTH, HEIGHT)
	gradient.AddColorStop(0, color.White)
	gradient.AddColorStop(1, color.Black)
	context.SetFillStyle(gradient)
	context.DrawRectangle(0, 0, WIDTH, HEIGHT)
	context.Fill()
}

func drawSimplexNoiseBackground (context *gg.Context) {
	seed := int64(time.Now().UTC().UnixNano())
        // good seeds: 1517604155637716269
	noiseGen := noise.NewWithSeed(seed)
	
	log.Printf("Seed for noise: %d", seed)

	scale := 4

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			x_val := float64(scale * x)/WIDTH
			y_val := float64(scale * y)/HEIGHT

			//calc noise but limit range of values to reduce intensity
			n := (noiseGen.Eval2(x_val, y_val) + 1) * (215 / 2) + 80 / 2
			noiseVal := uint8(n)
			context.SetColor(color.RGBA{ noiseVal, noiseVal, noiseVal, 255})
			context.SetPixel(x, y)
		}
	}
}
