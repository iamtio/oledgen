package main

import (
	"image"
	"image/color"
	"image/draw"
)

// Bar chart
type Bar struct {
	width, height, border        int
	positiveColor, negativeColor color.RGBA
	img                          *image.RGBA
}

// GetBar render bar with value from 0 to 1
func (b *Bar) GetBar(value float64, vertical bool) (*image.RGBA, error) {
	draw.Draw(b.img, b.img.Bounds(), &image.Uniform{b.positiveColor}, image.ZP, draw.Src)
	var cutFrom, cutTo image.Point

	if vertical {
		cutFrom = image.Point{b.border, b.border}
		cutTo = image.Point{b.width - b.border, b.height - int((float64(b.height-b.border*2) * value)) - b.border}
	} else {
		cutFrom = image.Point{b.border + int((float64(b.width-b.border*2) * value)), b.border}
		cutTo = image.Point{b.width - b.border, b.height - b.border}
	}
	draw.Draw(b.img, image.Rectangle{cutFrom, cutTo}, &image.Uniform{b.negativeColor}, image.ZP, draw.Src)
	return b.img, nil
}

// NewBar is constructor for Bar struct
func NewBar(width, height int, positiveColor color.RGBA) *Bar {
	negativeColor := color.RGBA{0, 0, 0, 0}
	return &Bar{width, height, 1, positiveColor, negativeColor, image.NewRGBA(image.Rect(0, 0, width, height))}
}
