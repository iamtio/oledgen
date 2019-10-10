package main

import (
	"image"
	"image/color"
	"image/draw"
)

// Bar chart
type Bar struct {
	width, height                int
	positiveColor, negativeColor color.RGBA
	img                          *image.RGBA
}

// GetBar render bar with value from 0 to 1
func (b *Bar) GetBar(value float64, vertical bool) (*image.RGBA, error) {
	draw.Draw(b.img, b.img.Bounds(), &image.Uniform{b.positiveColor}, image.ZP, draw.Src)
	var cutFrom, cutTo image.Point

	if vertical {
		cutFrom = image.Point{1, 1}
		cutTo = image.Point{b.width - 1, b.height - int((float64(b.height) * value)) - 1}
	} else {
		cutFrom = image.Point{1 + int((float64(b.width) * value)), 1}
		cutTo = image.Point{b.width - 1, b.height - 1}
	}
	draw.Draw(b.img, image.Rectangle{cutFrom, cutTo}, &image.Uniform{b.negativeColor}, image.ZP, draw.Src)
	return b.img, nil
}

// NewBar is constructor for Bar struct
func NewBar(width, height int, positiveColor color.RGBA) *Bar {
	negativeColor := color.RGBA{0, 0, 0, 0}
	return &Bar{width, height, positiveColor, negativeColor, image.NewRGBA(image.Rect(0, 0, width, height))}
}
