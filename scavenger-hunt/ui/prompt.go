package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Prompt struct {
	Container fyne.CanvasObject
	input     *widget.Entry
	button    *widget.Button
	onSubmit  func(string) // Callback for when input is submitted
}

func NewPrompt() *Prompt {
	p := &Prompt{
		input:  widget.NewMultiLineEntry(),
		button: widget.NewButton("Send", nil),
	}

	p.input.SetPlaceHolder("Your epic conversation starts here...")
	p.input.Wrapping = fyne.TextWrapWord

	// Set up button click handler
	p.button.OnTapped = func() {
		if p.onSubmit != nil {
			text := p.input.Text
			p.onSubmit(text)
			p.input.SetText("") // Clear input after sending
		}
	}

	buttonContainer := container.NewHBox(layout.NewSpacer(), p.button)
	p.Container = container.NewBorder(
		nil,
		buttonContainer,
		nil,
		nil,
		p.input,
	)

	return p
}

// SetOnSubmit sets the callback for when text is submitted
func (p *Prompt) SetOnSubmit(callback func(string)) {
	p.onSubmit = callback
}
