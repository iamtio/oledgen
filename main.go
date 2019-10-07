package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/tarm/serial"
)

var serialPort string
var seriallBaud int
var runMode string
var sleepTime time.Duration
var imageWidth int
var imageHeight int

func init() {
	flag.StringVar(&serialPort, "port", "", "Serial port connects to")
	flag.IntVar(&seriallBaud, "baud", 115200, "Serail port baudrate")
	flag.StringVar(&runMode, "mode", "", "Run mode: ascii, image, serial")
	flag.IntVar(&imageWidth, "width", 128, "Image width")
	flag.IntVar(&imageHeight, "height", 64, "Image height")
	flag.DurationVar(&sleepTime, "sleep", 1*time.Second, "Sleep between sendings")
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

	drawText(img, 1, fmt.Sprintf("ram:%7.2f/%.2f", float64(vMem.Used)/1024.0/1024.0/1024.0, float64(vMem.Total)/1024.0/1024.0/1024.0))
	for i := 0; i < len(cpus) && i < 4; i++ { // Limit to 4 cpus due to screen size
		drawText(img, 2+i, fmt.Sprintf("cpu%d:%6.2f %%", i+1, cpus[i]))
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
	for bit := uint(0); bit < 8; bit++ {
		fmt.Print((b >> bit) & 1)
	}
}
func printBlob(blob []uint8, bytesLineSize int) {
	for index, b := range blob {
		if index > 0 && index%bytesLineSize == 0 {
			fmt.Printf("\n")
		}
		printByte(b)
	}
}

type dummyWriter struct{}

func (dummyWriter) Write(b []byte) (int, error) {
	log.Printf("successfuly wrote: %d bytes\n", len(b))
	return len(b), nil
}
func (dummyWriter) Close() error {
	log.Printf("Closed dummy writer")
	return nil
}
func getDummyWriter() io.WriteCloser {
	return &dummyWriter{}
}
func getSerialWriter() io.WriteCloser {
	c := &serial.Config{Name: serialPort, Baud: seriallBaud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func main() {
	flag.Parse()
	switch runMode {
	case "ascii":
		img := generateImage(true, imageWidth, imageHeight)
		blob := getBlob(img, color.RGBA{0, 0, 0, 255})
		printBlob(blob, 16)
		fmt.Print("\n")
	case "image":
		img := generateImage(true, imageWidth, imageHeight)
		writeToFile(img, "out.png")
	case "serial":
		writer := getSerialWriter()
		defer writer.Close()
		for {
			img := generateImage(false, imageWidth, imageHeight)
			blob := getBlob(img, color.RGBA{0, 0, 0, 255})
			_, err := writer.Write(blob)
			if err != nil {
				panic(err)
			}
			time.Sleep(sleepTime)
		}
	default:
		fmt.Println("Nothing to do! write -mode flag to run something!")
		os.Exit(1)
	}

}
