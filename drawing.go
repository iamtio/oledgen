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
	cpusTotal, err := cpu.Percent(measureDuration, false)
	if err != nil {
		panic(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// RAM sprite
	draw.Draw(img, image.Rectangle{image.Point{0, 0}, image.Point{16, 16}}, sprites[1], image.ZP, draw.Src)
	// RAM bar
	bar := NewBar(74, 9, color.RGBA{0, 0, 0, 255})
	barImg, _ := bar.GetBar(vMem.UsedPercent*0.01, false)
	draw.Draw(img, image.Rectangle{image.Point{17, 3}, image.Point{width, height}}, barImg, image.ZP, draw.Src)

	// RAM text
	drawText(img, 1, fmt.Sprintf("%18.2f", float64(vMem.Used)/1024.0/1024.0/1024.0))
	// Cpu icon
	draw.Draw(img, image.Rectangle{image.Point{0, 18}, image.Point{width, height}}, sprites[0], image.ZP, draw.Src)
	// Cpu bars
	cpuBarWidth := 111 / len(cpus)
	for i := 0; i < len(cpus) && i < 12; i++ { // Limit to 12 cpus due to screen size
		bar := NewBar(cpuBarWidth, 30, color.RGBA{0, 0, 0, 255})
		barImg, _ := bar.GetBar(cpus[i]*0.01, true)
		draw.Draw(img, image.Rectangle{image.Point{17 + i*(cpuBarWidth-1), 18}, image.Point{width, height}}, barImg, image.ZP, draw.Src)
		// drawText(img, 2+i, fmt.Sprintf("   %d:%6.2f %%", i+1, cpus[i]))
	}
	drawText(img, 5, fmt.Sprintf("%8.2f%%", cpusTotal[0]))
	// stats, _ := disk.IOCounters("/dev/dm-0")
	// for k, v := range stats {
	// 	fmt.Printf("%s => %s", k, v)
	// }
	return img
}
