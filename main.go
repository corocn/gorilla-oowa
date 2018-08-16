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
	file := "https://emojipedia-us.s3.amazonaws.com/thumbs/320/twitter/147/gorilla_1f98d.png"
	pix, err = ansimage.NewScaledFromURL(
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
