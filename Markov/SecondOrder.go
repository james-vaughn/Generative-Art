package main

import (
	"image/color"
	"math/rand"
)

type SecondOrderMarkovChain struct {
	Chain     map[color.Color]map[color.Color]map[color.Color]int
	currColor color.Color
	currNeighbor color.Color
}

func NewSecondOrderChain() *SecondOrderMarkovChain {
	return &SecondOrderMarkovChain{
		make(map[color.Color]map[color.Color]map[color.Color]int),
		nil,
nil}
}

func (m *SecondOrderMarkovChain) Add(current, neighborInner, neighborOuter color.Color) {
	if m.currColor == nil {
		m.currColor = current
		m.currNeighbor = neighborInner
	}

	if m.Chain[current] == nil {
		m.Chain[current] = make(map[color.Color]map[color.Color]int)
	}

	if m.Chain[current][neighborInner] == nil {
		m.Chain[current][neighborInner] = make(map[color.Color]int)
	}


	m.Chain[current][neighborInner][neighborOuter] += 1
}

func (m *SecondOrderMarkovChain) Next() color.Color {
	var sum int

	for _, freq := range m.Chain[m.currColor][m.currNeighbor] {
		sum += freq
	}

	x := rand.Intn(sum)

	for color, freq := range m.Chain[m.currColor][m.currNeighbor] {
		//fmt.Println(color, freq)
		x -= freq

		if x <= 0 {
			m.currColor = m.currNeighbor
			m.currNeighbor = color
			return color
		}
	}

	return m.currColor
}
