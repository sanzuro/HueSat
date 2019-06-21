package main

import (
	"image/color"
	"math"
)

func toHSL(c color.Color) (float64, float64, float64) {
	var h, s, l float64
	R1, G1, B1, _ := c.RGBA()
	r := float64(uint8(R1)) / 255
	g := float64(uint8(G1)) / 255
	b := float64(uint8(B1)) / 255

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	l = (max + min) / 2

	delta := max - min
	if delta == 0 {

		return 0, 0, l
	}

	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	r2 := (((max - r) / 6) + (delta / 2)) / delta
	g2 := (((max - g) / 6) + (delta / 2)) / delta
	b2 := (((max - b) / 6) + (delta / 2)) / delta
	switch {
	case r == max:
		h = b2 - g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}

	switch {
	case h < 0:
		h++
	case h > 1:
		h--
	}

	return h, s, l
}

func toRGB(h, s, l float64) (float64, float64, float64) {

	if s == 0 {
		return l, l, l
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := hueToRGB(v1, v2, h+(1.0/3.0))
	g := hueToRGB(v1, v2, h)
	b := hueToRGB(v1, v2, h-(1.0/3.0))

	return r, g, b
}

func hueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h++
	}
	if h > 1 {
		h--
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	}
	return v1
}
