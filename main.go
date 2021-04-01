package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	MaxIters  = 500
	Width     = 1200
	Height    = 1200
	X1        = -2
	Y1        = 1.5
	X2        = 1
	Y2        = -1.5
	ImageName = "mandelbrot.png"
)

func initImg() *image.RGBA {
	start := image.Point{0, 0}
	end := image.Point{Width, Height}

	return image.NewRGBA(image.Rectangle{start, end})
}

func saveImg(img *image.RGBA) {
	f, err := os.Create(ImageName)
	// TODO: return an error
	if err == nil {
		png.Encode(f, img)
		return
	}

	fmt.Println("Error creating image")
}

func plotPixel(img *image.RGBA, x int, y int, pcolor color.Color) {
	img.Set(x, y, pcolor)
}

func mandel(c complex128) int {
	z := complex(0, 0)

	iters := 0
	for iters < MaxIters && cmplx.Abs(z) <= 2 {
		z = z*z + c
		iters++
	}

	return iters
}

func colorValue(iters int) color.Color {
	//plotc := color.RGBA{0, cc, cc, 255}
	plen := len(palette.Plan9)
	pcol := uint8(int((float64(plen)/float64(MaxIters))*float64(iters)) % plen)

	return palette.Plan9[pcol]
}

func pixelToComplex(x int, y int) complex128 {
	r := X1 + (float64(x)/Width)*(X2-X1)
	i := Y2 + (float64(y)/Height)*(Y1-Y2)

	return complex(r, i)
}

func progress(x int) {
	if x%(Width/10) == 0 {
		fmt.Printf("%d%%\n", int((float32(x)/float32(Width))*100))
	}
}

func render() {
	img := initImg()

	for x := 0; x < Width; x++ {
		progress(x)
		for y := 0; y < Height; y++ {
			cnum := pixelToComplex(x, y)
			iters := mandel(cnum)
			pcolor := colorValue(iters)
			plotPixel(img, x, y, pcolor)
		}
	}

	saveImg(img)
}

func main() {
	render()
}
