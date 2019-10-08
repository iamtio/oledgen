package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
)

func getSprites() []*image.RGBA {
	const spriteSize int = 16
	data, err := Asset("sprites.png")

	if err != nil {
		panic("Can't find sprite.png data in asset")
	}
	allSprites, err := png.Decode(bytes.NewReader(data))
	sprites := make([]*image.RGBA, 0, 0)
	for i := 0; i*spriteSize < allSprites.Bounds().Max.Y; i++ {
		s := image.NewRGBA(image.Rect(0, 0, spriteSize, spriteSize))
		draw.Draw(s, s.Bounds(), allSprites, image.Point{0, i*spriteSize - 1}, draw.Src)
		sprites = append(sprites, s)
		fmt.Println(i * spriteSize)
	}
	return sprites
}
