package ui

import (
	"image/color"
	"pixl/swatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildSwatches(app *AppInit) *fyne.Container { //the container type which is used to contain layouts so that we can place swatches into the layout and returning them into the container
	canvasSwatches := make([]fyne.CanvasObject, 0, 64) // this canvas swatches var acts as a buffer of canvas objects, so we just set the initial len 0 and cap 64
	// loop that builds the swatches
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255} // set initial color to white
		// create the swatch - in the for loop the code will remove borders and highlights around the swatches
		s := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				// deselect all the swatches
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			// we're going to select this swatch
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColor = s.Color
		})
		// initialization code - first for the zero index swatch, we're setting it to be selected. When a user opens the program the very forst swatch will be the selected swatch
		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0
			s.Refresh()
		}
		// append our swatches to the existing swatch slice
		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}
	// creating the grid wrap type of layout from fyne
	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)

}
