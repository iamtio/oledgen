package main

import "image"

func loadIcons() []*image.RGBA {
	cpuIcon := image.NewRGBA(image.Rect(0, 0, 16, 16))
	return []*image.RGBA{cpuIcon}
}
