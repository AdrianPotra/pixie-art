package main

import (
	"image/color"
	"pixl/apptype"
	"pixl/swatch"
	"pixl/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	// creating the app
	pixlApp := app.New()
	// creating the window
	pixlWindow := pixlApp.NewWindow("pixl")

	// creating app state
	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255}, //white
		SwatchSelected: 0,
	}

	// application initialization structure
	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}
	// we call the setup func that we just wrote
	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun() // part of toolkit and this will display the app and run it

}
