package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/james-vaughn/Generative-Art/SEIR/seir"
)

const (
	WIDTH       = 1000
	HEIGHT      = 1000
	SEED_OFFSET = WIDTH / 5
	N           = HEIGHT + SEED_OFFSET
)

func main() {
	outputImage := image.NewRGBA64(image.Rect(0, 0, WIDTH, HEIGHT))

	for i := 0; i < HEIGHT; i++ {
		model := modelFromSeed(i)

		for j := 0; j < WIDTH; j++ {
			color := seirPointToColor(model, j)
			outputImage.Set(i, j, color)
		}
	}

	f, err := os.Create("image.png")
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

func modelFromSeed(seed int) seir.SeirModel {
	s := float64((HEIGHT + SEED_OFFSET - seed) - SEED_OFFSET/2)
	config := seir.SeirConfig{
		Beta:  .9,
		Gamma: .1,
		Sigma: .3,
		Mu:    0,
		Nu:    0,
		S:     s + SEED_OFFSET,
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
	g := model.I[n]
	b := model.R[n]

	total := float64(r + g + b)
	rFrac := float64(r) / total
	gFrac := float64(g) / total
	bFrac := float64(b) / total

	return color.RGBA64{
		R: uint16(maxUint16 * rFrac),
		G: uint16(maxUint16 * gFrac),
		B: uint16(maxUint16 * bFrac),
		A: maxUint16,
	}
}
