package main

import (
	"image"
	"image/color" //needed to decode png
	"log"
	"math/rand"
	"time"

	"github.com/james-vaughn/Generative-Art/SEIR/seir"
	"github.com/james-vaughn/Generative-Art/Shared"
	noise "github.com/ojrac/opensimplex-go"
)

const (
	WIDTH        = 1920
	HEIGHT       = 1080
	N            = HEIGHT
	BOUNDARY     = .5 * HEIGHT
	FREQ_FACTOR  = .008
	SCALE_FACTOR = .25
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	outputImage := image.NewRGBA64(image.Rect(0, 0, WIDTH, HEIGHT))

	noiseGenerator := noise.NewWithSeed(rand.Int63())

	grid := make([]int, WIDTH*HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		seed := SCALE_FACTOR*HEIGHT*noiseGenerator.Eval2(float64(i)*FREQ_FACTOR, 0) + HEIGHT
		seirModel := modelFromSeed(seed)

		for j := 0; j < WIDTH; j++ {
			val := seirPointToVal(seirModel, j)
			grid[i*WIDTH+j] = val
		}
	}

	drawBackground(outputImage)
	drawSand(outputImage, grid, sandChain())
	drawImageGradient(outputImage, grid)
	Shared.SaveImage(outputImage, "position.png")
}

func sandChain() *Shared.MarkovChain {
	images := make([]image.Image, 0)
	imgList := []string{
		"input/guh.png",
	}

	for _, imgURL := range imgList {
		img, err := Shared.OpenImage(imgURL)

		if err != nil {
			log.Fatalf("Could not open image: %v", err)
		}

		images = append(images, img)
	}

	return Shared.MakeChain(images)
}

func drawImageGradient(outputImage *image.RGBA64, grid []int) {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			val := grid[i*WIDTH+j]

			currColor := outputImage.At(j, i).(color.RGBA64)

			if val == 0 {
				//todo fix blend
				c := color.RGBA64{0, 30000, 60000, 55000 - uint16(j*10)}
				outputImage.Set(j, i, Shared.Blend(currColor, c))
			}

			if val == 2 {
				c := color.RGBA64{0, 10000, 60000, 60000}
				outputImage.Set(j, i, Shared.Blend(currColor, c))
			}
		}
	}
}

func drawSand(outputImage *image.RGBA64, grid []int, markovChain *Shared.MarkovChain) {

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			val := grid[i*WIDTH+j]

			currColor := outputImage.At(j, i).(color.RGBA64)

			if val == 1 {
				r, g, b, _ := markovChain.Next().RGBA()
				sandColor := color.RGBA64{
					uint16(r),
					uint16(g),
					uint16(b),
					uint16(62000),
				}

				outputImage.Set(j, i, Shared.Blend(currColor, sandColor))
			}

			if val == 0 {
				r, g, b, _ := markovChain.Next().RGBA()
				sandColor := color.RGBA64{
					uint16(r),
					uint16(g),
					uint16(b),
					uint16(50000),
				}

				outputImage.Set(j, i, Shared.Blend(currColor, sandColor))
			}
		}
	}
}

func drawBackground(outputImage *image.RGBA64) {
	seed := int64(time.Now().UTC().UnixNano())
	// good seeds: 1517604155637716269
	noiseGen := noise.NewWithSeed(seed)

	scale := 3

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			x_val := float64(scale*x) / WIDTH
			y_val := float64(scale*y) / HEIGHT

			//calc noise but limit range of values to reduce intensity
			n := (noiseGen.Eval2(x_val, y_val)+1)*(215/2) + 80/2
			noiseVal := uint8(n)
			outputImage.Set(x, y, color.RGBA{noiseVal, noiseVal, noiseVal, 255})
		}
	}
}

func modelFromSeed(s float64) seir.SeirModel {
	config := seir.SeirConfig{
		Beta:  .9,
		Gamma: .1,
		Sigma: .3,
		Mu:    0,
		Nu:    0,
		S:     s,
		E:     0,
		I:     2,
		R:     N - s,
		Days:  40,
		Steps: WIDTH,
	}

	model := seir.NewModel(config)
	model.CalculatePoints()

	return model
}

func seirPointToVal(model seir.SeirModel, n int) int {
	const maxUint16 = 65535

	if model.S[n] > BOUNDARY {
		return 1
	}
	//g := model.I[n]
	if model.R[n] > BOUNDARY {
		return 2
	}

	return 0
}
