package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/james-vaughn/Generative-Art/SEIR/seir"
	"github.com/james-vaughn/Generative-Art/Shared"
)

const (
	WIDTH       = 1920
	HEIGHT      = 1080
	SEED_OFFSET = WIDTH / 5
	N           = HEIGHT + SEED_OFFSET
	NUM_CENTERS = 6
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	outputImage := image.NewRGBA64(image.Rect(0, 0, WIDTH, HEIGHT))

	colors := make([]color.Color, WIDTH*HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		model := modelFromSeed(i)

		for j := 0; j < WIDTH; j++ {
			color := seirPointToColor(model, j)
			colors[i*WIDTH+j] = color
		}
	}

	drawImageGradient(outputImage, colors)
	Shared.SaveImage(outputImage, "gradient.png")

	drawImageCircles(outputImage, colors)
	Shared.SaveImage(outputImage, "circles.png")

}

func drawImageGradient(outputImage *image.RGBA64, colors []color.Color) {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			color := colors[i*WIDTH+j]
			outputImage.Set(j, i, color)
		}
	}
}

func drawImageCircles(outputImage *image.RGBA64, colors []color.Color) {
	var countToColor int = 1
	colored := make(map[image.Point]bool)
	toColor := make([]image.Point, NUM_CENTERS)
	colorCounter := 0
	for i := 0; i < len(toColor); i++ {
		toColor[i] = image.Point{rand.Intn(WIDTH), rand.Intn(HEIGHT)}
	}

coloringLoop:
	for countToColor > 0 {
		idx := rand.Intn(countToColor)
		pt := toColor[idx]
		toColor[idx] = toColor[len(toColor)-1] // Replace it with the last one.
		toColor = toColor[:len(toColor)-1]

		countToColor--

		//if pt.X < 0 || pt.X >= WIDTH || pt.Y < 0 || pt.Y >= HEIGHT {
		//	continue coloringLoop
		//}

		//prevent duplicated
		if colored[pt] == true {
			continue coloringLoop
		}

		color := colors[colorCounter] //next
		colorCounter += 1

		outputImage.Set(pt.X, pt.Y, color)

		//context.SetColor(toRGBA(c))
		//context.SetPixel(x, y)
		colored[pt] = true

		if pt.X > 0 && pt.X < WIDTH-1 && pt.Y > 0 && pt.Y < HEIGHT-1 {
			//append neighbors
			toColor = append(toColor,
				image.Point{pt.X - 1, pt.Y},
				image.Point{pt.X + 1, pt.Y},
				image.Point{pt.X, pt.Y - 1},
				image.Point{pt.X, pt.Y + 1})

			countToColor += 4
		}
	}
}

func modelFromSeed(seed int) seir.SeirModel {
	s := float64(HEIGHT + SEED_OFFSET - seed)
	config := seir.SeirConfig{
		Beta:  .9,
		Gamma: .1,
		Sigma: .3,
		Mu:    0,
		Nu:    0,
		S:     s + SEED_OFFSET,
		E:     0,
		I:     1,
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
