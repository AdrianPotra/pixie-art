package ui

// creating a setup function for the UI elements
func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	//access the app, more specifically the window
	app.PixlWindow.SetContent(swatchesContainer)
}
