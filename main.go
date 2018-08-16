package main

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/eliukblau/pixterm/ansimage"
	"image/color"
)

func main() {
	drawAsciiArtGorilla()
	println("Oowa Oowa !!")
}

func drawTwemojiGorilla() {
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

	if err == nil {
		pix.Draw()
	}
}

func drawAsciiArtGorilla() {
	doc := heredoc.Doc(`
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMHkkbkkkkkkkkkkbbkkkHMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMkkkkkkkkkbkkkkkkkkkkkbkkkkkkkqMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkbkkkMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMkkkkkkkkkkkbkkkkkkkkkkkkkkkkkkbkkkkkbkkkbHMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMkkkkkkkkkkkkkkkkbkkbkkbkkbkkbkkkkkkkkkkkkkkkkkHMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMkkkkkkkkbkkbkkkkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkkkkHMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMkkkkkkkkkkkkkkkkbkkkkkkkkkkkkkkkkkkbkkkkkkbkkbkkkkkkkkMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMkkkkkkkkkkkkkkkkkkkkbkkbkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkkHMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMkkkkkkkkkkkkbkkbkkkkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkbkkkkkkkkMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMkkkkkkkkkkkkkkkkkkkbkkkkkkkkkkkkkkkkkkbkkkkkkbkkbkkkkkkkkkkkkkMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMkkkkkkkkkkkkkkkkkkkkkkkbkkbkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkkkkkkkMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMkkkkkkkkkkkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkbkkbkkkkkkkkkMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkkkkkkkkkbkkkkkkbkkbkkkkkkkkkkkkkkkkkMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMkkkkkkkkkbkkkkkkkkkkkkkkkkbkkbkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMHkkkkkkkkkkkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkbkkbkkkkkkkkkkkkHMMMMMMMMMMMMMM
MMMMMMMMMMMMMNkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkkkkkkkkkkkbkkkkkbkkbkkkkkkkkkkkkkkkkkkkkHMMMMMMMMMMMMM
MMMMMMMMMMMMMkkkkkkkkbkkbkkkkkkkkkkkkkkbkkbkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkkbkkbkkkkkkkkMMMMMMMMMMMMM
MMMMMMMMMMMMMkkkkbkkkkkkkkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkbkkbkkkkkkkkkbkkkkkMMMMMMMMMMMMM
MMMMMMMMHHMMHkkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkkkkkkkkkbkkkkkkkbkkbkkkkkkkkkkkkkkkkkkkkkHMMHHMMMMMMMM
MMMMHkk&+>zUbkkkkkkkbkkbkkkkkkkkkkkkkkkbkkbkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkbkkbkkkkkkkbkk9C>+&dkHMMMM
MMMHkkkkkk>>?HkkbkkkkkkkkkbkkbkkbkkkkkkkkkkkkkkkkkkkkkkkkbkkkkkkkkkkbkkbkkkkkkkkbkkkkkH3>>dkkkkkHMMM
MM#bkkkkkkI>>dkkkkkkkkHWUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUWHkkkkkkkk0>>jkkkkkkkMMM
MM#kkkkkkkR>+dkkkkkkH6>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>;>>>>>>>vHkkkkkkk+;dkkkkkbkMMM
MMNkkkkkkbHQWkkkkbkk0>>;>>>>;>>+&&&&&&&&&<>>>>>>>>>>>;>;>;>j&&&&&&&&+>>;>>;>;>>dkkbkkkbHkWkkkkkkkMMM
MMMNkkkkkkkbkkkkkkkkI>>>>;>>>>+XkkkkkkkkkR>>;>;>;>;>>;>>;>dkbkkkkkkkHz>>;>>;>;>jkkkkkkkkkkkkbkkkHMMM
MMMNkkkkkkkkkkkkkkkkS>>;>>;>>>+WkkkkkkkkkD>>>>>>>>>>>>;>>>dkkkkkkkkkHz>;>>;>>>;dkkkkkkkkkkkkkkkkMMMM
MMMMNkkkkkkkkkbkkkkkR<>>>;>>;>>vHkkkkkkkK1>>;>>;>>;>>;>>;>?WkkkbkkkH6>>>;>>;>>+dkkkbkkkkkkkkkkkNMMMM
MMMMMNkkkkkbkkkkkkkkHI>;>>>>>>>>+7UWHU9C>>>>>;>>>>>;>>>;>>>>?TWHW96<>>;>>;>>;>jWkkkkkkbkkbkkkkNMMMMM
MMMMMMNHkkkkkkkkkbkkkR>>>;>>;>>>>>>>>>>>>>;>>>>>;>>>>;>>>>>;>>>>>>>>>;>>;>>;>>dkkkkkkkkkkkkkHNMMMMMM
MMMMMMMNHkkkkkbkkkkkkk2>>;>>>;>>>>>>;<<!~'   '~<<>><<!''  '_!<;>>;>>;>>;>>;>>jkkkkkbkkkkkkkNMMMMMMMM
MMMMMMMMMNkkkkkkkkkkkkR+>>>>>>;>;>><'             '            _<>;>>;>>>;>>+dkkkkkkkkbkkkHMMMMMMMMM
MMMMMMMMMHkkkkkkkbkkkkkk>>;>>>>>>;:     ..JJJ-..    ...JJJ-.     (>;>>;>>;>>dkkkkkkkkkkkkkHMMMMMMMMM
MMMMMMMMMHkkkbkkkkkkkkkf>>>;>;>>>;_   .jWkkkkkkc><(>+XkkbkkHc-    >>>;>>;>>>ZkkkkkbkkkkkkkHMMMMMMMMM
MMMMMMMMMNkkkkkkkkkkkkK<>;>>>>;>>><...>zHkkkkkWz>>>>?WkkkkkH3>_..(>>;>>;>>;><4kkkkkkkbkkkkNMMMMMMMMM
MMMMMMMMMNkkkkkkbkkkbHC>>;>>>>>>>>>>>>>>>?TTT1>>>>>>>>?7TTC<>>;;>>;>>;>>>;>>;jHkkkkkkkkkkkMMMMMMMMMM
MMMMMMMMMNHkkbkkkkkkkD>;>>>>;>>;>>>;>;>>>>>>>>>>>>>>>;>>>>>>>>;>;>>;>>;>>;>>>>dkbkkkkkkkkHMMMMMMMMMM
MMMMMMMMMMkkkkkkkkkkHC>>>;>>>>>>;>>>>>>>>>>;>>>>>;>>>>>>;>>;>>>>>;>>;>>;>>;>;>zbkkkbkkbkkkMMMMMMMMMM
MMMMMMMMMMHkkkkkbkkkHz>;>>>;>>;>>>;>>;>>;>>>>>;>>>;>>;>;>;>>;;>;>>;>>;>>;>>;>;jWkkkkkkkkkNMMMMMMMMMM
MMMMMMMMMMNkkkkkkkkkH<>;>>>>>>>>>>>>>>;>>>;>>>>>>>>>;>>>;>>;>>;>>;>>;>>;>>>>>>jWkkkkkkkkkMMMMMMMMMMM
MMMMMMMMMMMNkkkkkkkkHz>>>;>>>>>>>;>>;>>>>>>;>>;>>;>>>;>>>>;>>>>>;>>>>>+>>>;>;>jkkkbkkbkkHMMMMMMMMMMM
MMMMMMMMMMMNHkkkbkkkkZ>>>+dkkkkkbbWkkmmAAaaaa&&&&&&&&&&aaaaQAAQkkWbbkkkkkR<>>+dkkkkkkkkHMMMMMMMMMMMM
MMMMMMMMMMMMNkkkkkkkkHc;>>+OTWHkkkkkkkbbkkkkkkkkkkkkkkkkkkkkkkkkkkkkkHW96<>>>jkbkkkkkkkMMMMMMMMMMMMM
MMMMMMMMMMMMMNkkkkkkkkR+>>>>>>><17TTUUUUWWHHHHHHHHHHHHHHHHHWUU9UTTC1<>>>>>>>jbkkkkbkkkMMMMMMMMMMMMMM
MMMMMMMMMMMMMMNkkkkbkkkHx>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>;>>;>>;>>>>>>>>>>;+dbkkkkkkkkNMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMNHkkkkkbbkmx>>>>>>>>;>;>>;>>>>>>>>>>>>;>>;>>;>>;>;>;>>;>>+dWkkkkkkkkHMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMNHkbkkkkkkmx+>;>>;>>>>>>>;>;>>;>>;>>;>>>;>>;>>;>>;>;+&dbkkkkkkkkHNMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMNHkkkkkkkkbm&+>>>>>;>>>>>>;>>>>>>>>;>;>>;>>>>>>+udWkkkkkkkbkkNMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMNNkkkkkkkkkkHme&+>>;>>;>>>>;>>;>>>;>>;>++ugQWkkkkkkkkkkkNMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMNNHHkkkkkkkkkkkWkmA&&&&++++++&&&gQkWkkkkkkkkkkkkHHNMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNNHHkkkkkkkkkkkkkkkkkkkbkkkkkkkkkkkqHHNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNNNNNNNNNNNNNNNNNNNNNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM`)
	println(doc)
}
