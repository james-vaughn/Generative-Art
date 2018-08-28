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

	colors := make([]color.Color, WIDTH*HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		seed := SCALE_FACTOR*HEIGHT*noiseGenerator.Eval2(float64(i)*FREQ_FACTOR, 0) + HEIGHT
		seirModel := modelFromSeed(seed)

		for j := 0; j < WIDTH; j++ {
			color := seirPointToColor(seirModel, j)
			colors[i*WIDTH+j] = color
		}
	}

	drawImageGradient(outputImage, colors)
	drawSand(outputImage, sandChain())
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

func drawImageGradient(outputImage *image.RGBA64, colors []color.Color) {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			color := colors[i*WIDTH+j]
			outputImage.Set(j, i, color)
		}
	}
}

func drawSand(outputImage *image.RGBA64, markovChain *Shared.MarkovChain) {

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			r, _, _, _ := outputImage.At(j, i).RGBA()

			if r > 0 {
				outputImage.Set(j, i, markovChain.Next())
			}
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

func seirPointToColor(model seir.SeirModel, n int) color.Color {
	const maxUint16 = 65535

	r := model.S[n]
	//g := model.I[n]
	b := model.R[n]

	if r > BOUNDARY {
		return color.RGBA64{65535, 0, 0, 65535}
	}

	if b > BOUNDARY {
		return color.RGBA64{0, 0, 65535, 65535}
	}

	return color.RGBA64{0, 65535, 0, 65535}
	//total := float64(r + g + b)
	//rFrac := float64(r) / total
	//gFrac := float64(g) / total
	//bFrac := float64(b) / total
	//
	//return color.RGBA64{
	//	R: uint16(maxUint16 * rFrac),
	//	G: uint16(maxUint16 * gFrac),
	//	B: uint16(maxUint16 * bFrac),
	//	A: maxUint16,
	//}
}
