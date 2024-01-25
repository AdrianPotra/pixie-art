/*
Author: Adrian Potra
Version 1.0

Notes: we will import the fyne package "fyne.io/fyne/v2" - more info on https://github.com/fyne-io/fyne
*/

package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType int

// creating a canvas config, which will store some info about the canvas
type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int // represent the nr of rows and columns in the pixel art
	PxSize         int // will represent the scale factor of the pixels - visually you can notice how many pixels are on screen

}

// general application state
type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int // swatches will be stored in a slice, so this int is basically an index of the slice
	FilePath       string
}

// receiver function to set the file path
func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
