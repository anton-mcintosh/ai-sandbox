package ui

import (
	"hunt/types"
	_ "image/color"

	"fyne.io/fyne/v2"
	_ "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Think struct {
	Container fyne.CanvasObject
	textArea  *widget.Entry
	thinker   types.ThinkUpdater
	// Add other fields needed
}

func NewThink(thinker types.ThinkUpdater) *Think {
	t := &Think{
		textArea: widget.NewMultiLineEntry(),
		thinker:  thinker,
	}
	t.textArea.Wrapping = fyne.TextWrapWord
	t.textArea.MultiLine = true

	t.textArea.MinSize()
	scroll := container.NewScroll(t.textArea)
	t.Container = container.NewStack(scroll)
	return t
}
func (t *Think) UpdateThinking(content string) {
	t.textArea.SetText(t.textArea.Text + content)
	t.textArea.CursorRow = len(t.textArea.Text)
	t.textArea.Refresh()
}
