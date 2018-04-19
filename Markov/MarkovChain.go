package main

import (
	"image/color"
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

