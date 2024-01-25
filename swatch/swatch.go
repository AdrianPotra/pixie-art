/*
Author: Adrian Potra
Version 1.0

Notes: for swtaches we will be using the widget and widget renderer interfaces from the fyne package. Some basic info about both below:

type Widget interface {
	- base functionality and state for all widgets (size, position etc.)
	- initialized with widget.ExtendBaseWidget(widget)
	- it has 2 components:
	  1. CanvasObject - keeps track of sie, position etc. and also provides functionality to set and retrieve those
	  2. CreateRenderer() WidgetRenderer - create renderer function will be called by the toolkit itself and the return object widget renderer will define how it actually looks
  }

  type WidgetRenderer interface {
   - there are many more methods included in this interface, but we will use the following:
     Destroy() - to be left blank in this implementation
	 Layout(Size) - calculate the position of individual objects that make up this widget
	 MinSize() Size - min dimensions that the widget occupies
	 Objects[] CanvasObject - all objects that should be drawn
	 Refresh() - an update occured and the widget needs to be redrawn (by update it means user actions like mouse clicks or movements, resizes of window)
  }

}
*/

package swatch

import (
	"image/color"
	"pixl/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int             // used to select a swatch
	clickHandler func(s *Swatch) // we will supply code on  what to do when it's clicked upon - we want the color to change when we click on it basically
}

// receiver function to set the color of the swatch
func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

// helper function to create a new swatch
func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	// here we will set some defaults
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}
	swatch.ExtendBaseWidget(swatch) // we supply or current swatch to the extend base widget fyne package function
	return swatch
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  square,
		objects: objects,
		parent:  swatch,
	}
}
