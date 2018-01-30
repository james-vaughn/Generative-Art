package main

import (
	"image/color"
	"github.com/fogleman/gg"
)

func main() {
	context := gg.NewContext(1000,1000)

	context.SetColor(color.RGBA {255, 0, 0, 255})
	context.DrawRectangle(100, 100, 800, 800)
	context.Fill()
	context.SavePNG("../output/perlin.png")
}
