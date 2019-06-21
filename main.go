package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type kolor struct {
	r, g, b, a uint32
}

func main() {

	file, err := os.Open("photo.png")

	if err != nil {
		fmt.Print(" error occured on the first step ")
		os.Exit(0)
	}
	img, err := png.Decode(file)
	rect := img.Bounds()
	nimg := image.NewRGBA(rect)
	for y := 0; y < rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {
			// r, g, b, _ := img.At(x, y).RGBA()
			h, s, l := toHSL(img.At(x, y))
			h = 0.33
			r, g, b := toRGB(h, s, l)
			nimg.Set(x, y, color.NRGBA{
				R: uint8(r * 255),
				G: uint8(g * 255),
				B: uint8(b * 255),
				A: uint8(255),
			})
		}
	}
	w, _ := os.Create("kate.png")
	defer w.Close()
	_ = png.Encode(w, nimg)
}
