package main

import (
	"github.com/eliukblau/pixterm/ansimage"
	"image/color"
	"os"
)

func main() {
	var (
		pix *ansimage.ANSImage
		err error
	)

	// get terminal size
	tx, ty := 80, 24

	// set image scale factor for ANSIPixel grid
	sfy, sfx := 2, 1

	// load from file
	file := "./assets/gorilla.png"

	pix, err = ansimage.NewScaledFromFile(
		file,
		sfy*ty, sfx*tx,
		color.Black,
		ansimage.ScaleModeResize,
		ansimage.NoDithering)

	if err != nil {
		os.Exit(1)
	}

	pix.Draw()
	println("Oowa Oowa !!")
}
