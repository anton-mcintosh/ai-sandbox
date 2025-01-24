package ui

import (
	"hunt/types"
	_ "image/color"

	"fyne.io/fyne/v2"
	_ "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/* This is really only necessary in reasoning models such as the 
* DeepSeek-R1.
*/

type Think struct {
	Container fyne.CanvasObject
	textArea  *widget.Label
	thinker   types.ThinkUpdater
	scroll    *container.Scroll
}

func NewThink(thinker types.ThinkUpdater) *Think {
	t := &Think{
		textArea: widget.NewLabel(""),
		thinker:  thinker,
	}
	t.textArea.Wrapping = fyne.TextWrapWord

	t.textArea.MinSize()
	t.scroll = container.NewScroll(t.textArea)
	t.Container = container.NewStack(t.scroll)
	return t
}
func (t *Think) UpdateThinking(content string) {
	t.textArea.SetText(t.textArea.Text + content)
	t.textArea.Refresh()
	t.scroll.ScrollToBottom()
}
