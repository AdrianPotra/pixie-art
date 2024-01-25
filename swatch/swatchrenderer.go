/*
Author: Adrian Potra
Version 1.0
*/

package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type SwatchRenderer struct {
	square  *canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch // this points to the swatch that created this renderer - needed in order to determine whether or not something is selected in our refresh function
}

// interface implementation for the widget renderer
func (renderer *SwatchRenderer) MinSize() fyne.Size {
	return renderer.square.MinSize() // canvas.Rectangle is also a widget, so min size is defined for all widgets
}

// layout function - it will determine where in the layout the swatches will be placed
func (renderer *SwatchRenderer) Layout(size fyne.Size) {
	// since our swatch consists only of a single object in the canvas slice, we can access it directly - we also apply the resize func and supply the size - it will resize our existing square to the max available size
	renderer.objects[0].Resize(size)
}

// refresh function - we are going to update the internal state
func (renderer *SwatchRenderer) Resize() {
	renderer.Layout(fyne.NewSize(20, 20))             //we provide our own size of 20 for the swatch
	renderer.square.FillColor = renderer.parent.Color // set the color to whatever the parent happens to be
	// we check to see if the parent is selected
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3                               // we'll set it to 3 pixels
		renderer.square.StrokeColor = color.NRGBA{255, 255, 255, 255} // set the color to white
		// we need to update our existing render object
		renderer.objects[0] = renderer.square
	} else {
		//we remove the stroke and we reset the render object
		renderer.square.StrokeWidth = 3
		renderer.objects[0] = renderer.square
	}
	canvas.Refresh(renderer.parent)
}

// objects function that will return the objects already created
func (renderer *SwatchRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

// destroy function - it's empty, but it needs to be mentioned (required) in the implementation
func (renderer *SwatchRenderer) Destroy() {
}
func (renderer *SwatchRenderer) Refresh() {
}
