package ui

import (
	"pixl/apptype"
	"pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window // application window that we drop things into
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
