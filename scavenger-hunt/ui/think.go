package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Think struct {
	Container fyne.CanvasObject
	thinking  *canvas.Text
	// Add other fields needed
}

func NewThink() *Think {
	text := canvas.NewText("Thinking...", color.White)
	return &Think{
		thinking:  text,
		Container: text,
	}
}
