package main

import (
	"fmt"
	"image"
	_ "image/png" //needed to decode png
	"log"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/james-vaughn/Generative-Art/Shared"
)

const (
	WIDTH       = 1000
	HEIGHT      = 1000
	NUM_CENTERS = 10
)

//https://jonnoftw.github.io/2017/01/18/markov-chain-image-generation
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	images := make([]image.Image, 0)
	imgList := []string{
		"input/sand.png",
	}

	for _, imgURL := range imgList {
		img, err := Shared.OpenImage(imgURL)

		if err != nil {
			log.Fatalf("Could not open image: %v", err)
		}

		images = append(images, img)
	}

	fmt.Println("Making chain")
	markovChain := Shared.MakeChain(images)

	context := gg.NewContext(WIDTH, HEIGHT)
	fmt.Println("Drawing image")
	drawImage(context, markovChain)

	context.SavePNG("output/markov.png")
}

func drawImage(context *gg.Context, chain *Shared.MarkovChain) {
	var countToColor int = 1
	colored := make(map[image.Point]bool)
	toColor := make([]image.Point, NUM_CENTERS)

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

		c := chain.Next()

		context.SetColor(c)
		context.SetPixel(pt.X, pt.Y)

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

//func makeSecondOrderChain(image image.Image) *SecondOrderMarkovChain {
//	markovChain := NewSecondOrderChain()
//
//	minY := image.Bounds().Min.Y
//	minX := image.Bounds().Min.X
//	maxY := image.Bounds().Max.Y
//	maxX := image.Bounds().Max.X
//
//	for col := minY; col < maxY; col++ {
//		for row := minX; row < maxX; row++ {
//
//			//add neighbors
//			for dH := -1; dH <= 1; dH++ {
//				for dV := -1; dV <= 1; dV++ {
//					for dH_2 := -1; dH_2 <= 1; dH_2++ {
//						for dV_2 := -1; dV_2 <= 1; dV_2++ {
//							if dV != 0 && dH != 0 && dV_2 != 0 && dH_2 != 0 {
//								markovChain.Add(
//									image.At(clamp(row, minX, maxX-1), clamp(col, minY, maxY-1)),
//									image.At(clamp(row+dV, minX, maxX-1), clamp(col+dH, minY, maxY-1)),
//									image.At(clamp(row+dV+dV_2, minX, maxX-1), clamp(col+dH+dH_2, minY, maxY-1)))
//							}
//						}
//					}
//				}
//			}
//
//		}
//	}
//
//	return markovChain
//}
