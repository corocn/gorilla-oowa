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

	// load from file
	file := "./assets/gorilla.png"
	pix, err = ansimage.NewScaledFromFile(
		file,
		60, 76,
		color.Black,
		ansimage.ScaleModeResize,
		ansimage.NoDithering)

	if err != nil {
		os.Exit(1)
	}

	pix.Draw()
	println("Oowa Oowa !!")
}
