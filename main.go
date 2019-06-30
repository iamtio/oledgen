package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
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
	const height int = 11
	addLabel(img, 1, height*line, label)
}

func main() {
	v, _ := mem.VirtualMemory()
	ones, _ := time.ParseDuration("3s")
	cpus, _ := cpu.Percent(ones, true)

	img := image.NewRGBA(image.Rect(0, 0, 128, 64))

	drawText(img, 1, fmt.Sprintf("ram: %.2f/%.2f Gb", float64(v.Used)/1024.0/1024.0/1024.0, float64(v.Total)/1024.0/1024.0/1024.0))
	for i := 0; i < len(cpus); i++ {
		drawText(img, 2+i, fmt.Sprintf("cpu%d: %5.2f %%", i+1, cpus[i]))
	}

	f, err := os.Create("hello-go.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
