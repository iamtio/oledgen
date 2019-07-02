package main

import (
	"flag"
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

var serialPort string

func init() {
	flag.StringVar(&serialPort, "port", "", "Serial port connects to")
}
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
	const height int = 11
	addLabel(img, 1, height*line, label)
}

func generateImage(first bool) *image.RGBA {
	v, _ := mem.VirtualMemory()
	var ones time.Duration
	if first {
		ones, _ = time.ParseDuration("3s")
	}
	cpus, _ := cpu.Percent(ones, true)

	img := image.NewRGBA(image.Rect(0, 0, 128, 64))

	drawText(img, 1, fmt.Sprintf("ram: %.2f/%.2f Gb", float64(v.Used)/1024.0/1024.0/1024.0, float64(v.Total)/1024.0/1024.0/1024.0))
	for i := 0; i < len(cpus); i++ {
		drawText(img, 2+i, fmt.Sprintf("cpu%d: %5.2f %%", i+1, cpus[i]))
	}
	return img
}
func writeToFile(img *image.RGBA, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
func getBlob(img *image.RGBA, positiveColor color.RGBA) []uint8 {
	rect := img.Rect

	blobSize := rect.Max.X * rect.Max.Y / 8
	blob := make([]uint8, blobSize)

	var bitNum int
	for y := 0; y < rect.Max.Y; y++ {
		for x := 0; x < rect.Max.X; x++ {
			bitShift := uint(bitNum % 8)
			byteNum := bitNum / 8

			bit := uint8(0)
			if img.At(x, y) == positiveColor {
				bit = 1
			}
			blob[byteNum] |= (bit << bitShift)
			bitNum++
		}
	}
	return blob
}
func printByte(b uint8) {
	fmt.Printf("%d%d%d%d%d%d%d%d",
		(b>>0)&1,
		(b>>1)&1,
		(b>>2)&1,
		(b>>3)&1,
		(b>>4)&1,
		(b>>5)&1,
		(b>>6)&1,
		(b>>7)&1,
	)
}
func printBlob(blob []uint8, bytesLineSize int) {
	for index, b := range blob {
		if index%bytesLineSize == 0 && index != 0 {
			fmt.Printf("\n")
		}
		printByte(b)
	}
}
func main() {
	img := generateImage(true)
	blob := getBlob(img, color.RGBA{0, 0, 0, 255})
	printBlob(blob, 16)
	writeToFile(img, "hello-go.png")
}
