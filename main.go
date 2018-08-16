package main

import (
	"github.com/eliukblau/pixterm/ansimage"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {
	println("\U0001F98D < Oowa Oowa !!")

	var pix *ansimage.ANSImage

	// get terminal size
	tx := 80
	ty := 24

	// get scale mode from flag
	sm := ansimage.ScaleMode(0)

	// get dithering mode from flag
	dm := ansimage.DitheringMode(0)

	// set image scale factor for ANSIPixel grid
	sfy, sfx := 2, 1

	mc, _ := colorful.Hex("#000000") // RGB color from Hex format

	file := "./assets/gorilla.png"

	pix, _ = ansimage.NewScaledFromFile(file, sfy*ty, sfx*tx, mc, sm, dm)

	pix.Draw()
}
