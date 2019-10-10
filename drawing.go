package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func drawText(img *image.RGBA, line int, label string) {
	const height int = 12
	addLabel(img, 1, height*line, label)
}

func generateImage(first bool, width, height int) *image.RGBA {
	vMem, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	var measureDuration time.Duration
	if first {
		measureDuration = 1 * time.Second
	}
	cpus, err := cpu.Percent(measureDuration, true)
	if err != nil {
		panic(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Ram sprite
	draw.Draw(img, image.Rectangle{image.Point{0, 0}, image.Point{16, 16}}, sprites[1], image.ZP, draw.Src)
	// Ram text
	drawText(img, 1, fmt.Sprintf("%7.2f/%.2f", float64(vMem.Used)/1024.0/1024.0/1024.0, float64(vMem.Total)/1024.0/1024.0/1024.0))
	// Cpu icon
	draw.Draw(img, image.Rectangle{image.Point{0, 16}, image.Point{16, 32}}, sprites[0], image.ZP, draw.Src)

	for i := 0; i < len(cpus) && i < 4; i++ { // Limit to 4 cpus due to screen size
		drawText(img, 2+i, fmt.Sprintf("   %d:%6.2f %%", i+1, cpus[i]))
	}

	// stats, _ := disk.IOCounters("/dev/dm-0")
	// for k, v := range stats {
	// 	fmt.Printf("%s => %s", k, v)
	// }

	// bar := NewBar(8, 32, color.RGBA{0, 0, 0, 255})
	// barImg, _ := bar.GetBar(0.5, true)
	// draw.Draw(img, barImg.Bounds(), barImg, image.ZP, draw.Src)
	return img
}
