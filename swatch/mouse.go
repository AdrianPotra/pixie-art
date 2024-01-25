/*
Author: Adrian Potra
Version 1.0

Notes: we need to implement here Mouse interfaces. Details below:

type Mouseable interface {
	MouseDown(*MouseEvent) - mouse button pressed
	MouseUp(*MouseEvent) - mouse button released
}

type Scrollable interface {
	Scrolled(*ScrollEvent) - mouse scroll wheel movement
}
type Hoverable interface {
	MouseIn(*MouseEvent) - mouse pointer enters an element
	MouseMoved(*MouseEvent) - mouse pointer moved over an element
	MouseOut() - mouse pointer leaves an element
}

*/

package swatch

import "fyne.io/fyne/v2/driver/desktop"

// function for when user clicks down the mouse button - it's going to select the swatch
func (swatch *Swatch) MouseDown(ev *desktop.MouseEvent) {
	swatch.clickHandler(swatch)
	swatch.Selected = true
	swatch.Refresh()
}

func (swatch *Swatch) MouseUp(ev *desktop.MouseEvent) {

}
