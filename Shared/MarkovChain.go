package Shared

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

type MarkovChain struct {
	Chain     map[color.Color]map[color.Color]int
	currColor color.Color
}

func NewChain() *MarkovChain {
	return &MarkovChain{
		make(map[color.Color]map[color.Color]int),
		nil}
}

func (m *MarkovChain) Add(current, neighbor color.Color) {
	if m.currColor == nil {
		m.currColor = current
	}

	if m.Chain[current] == nil {
		m.Chain[current] = make(map[color.Color]int)
	}

	m.Chain[current][neighbor] += 1
}

func (m *MarkovChain) Next() color.Color {
	var sum int

	for _, freq := range m.Chain[m.currColor] {
		sum += freq
	}

	x := rand.Intn(sum)

	for color, freq := range m.Chain[m.currColor] {
		//fmt.Println(color, freq)
		x -= freq

		if x <= 0 {
			m.currColor = color
			//fmt.Println(m.Chain[color])
			return color
		}
	}

	return m.currColor
}

func toRGBA(c color.Color) color.Color {
	r, g, b, a := c.RGBA()

	return color.RGBA{
		uint8(r),
		uint8(g),
		uint8(b),
		uint8(a),
	}
}

func MakeChain(images []image.Image) *MarkovChain {
	markovChain := NewChain()

	for _, image := range images {
		minY := image.Bounds().Min.Y
		minX := image.Bounds().Min.X
		maxY := image.Bounds().Max.Y
		maxX := image.Bounds().Max.X

		for col := minY; col < maxY; col++ {
			for row := minX; row < maxX; row++ {

				//add neighbors
				for dH := -1; dH <= 1; dH++ {
					for dV := -1; dV <= 1; dV++ {
						if dV != 0 && dH != 0 {
							markovChain.Add(
								image.At(clamp(row, minX, maxX-1), clamp(col, minY, maxY-1)),
								image.At(clamp(row+dV, minX, maxX-1), clamp(col+dH, minY, maxY-1)))
						}
					}
				}

			}
		}
	}

	//connect the image chain nodes
	var prevColor color.Color
	for _, image := range images {
		if prevColor == nil {
			prevColor = image.At(0, 0)
			continue
		}

		currColor := image.At(0, 0)
		markovChain.Add(prevColor, currColor)
		markovChain.Add(currColor, prevColor)

		prevColor = currColor
	}

	return markovChain
}

func clamp(x, min, max int) int {
	return int(math.Max(float64(min), math.Min(float64(max), float64(x))))
}
