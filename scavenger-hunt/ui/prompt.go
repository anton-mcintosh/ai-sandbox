package ui

import (
	_ "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Prompt struct {
	Container fyne.CanvasObject
	input     *widget.Entry
	button    *widget.Button
	onSubmit  func(string) 
}

func NewPrompt() *Prompt {
	p := &Prompt{
		input:  widget.NewMultiLineEntry(),
		button: widget.NewButton("Send", nil),
	}

	p.input.SetPlaceHolder("Your epic conversation starts here...")
	p.input.Wrapping = fyne.TextWrapWord
	//after much frustration, this is how you capture the enter key for submission
	//but unfortunately it's shift+enter to submit if using the NewMultiLineEntry
	p.input.OnSubmitted = func(s string) {
		p.submit()
	}
	p.button.OnTapped = p.submit

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

func (p *Prompt) submit() {
	if p.onSubmit != nil && p.input.Text != "" {
		//this very specific ordering allows the input to be cleared before thinking commences.
		text := p.input.Text
		p.input.SetText("")
		p.onSubmit(text)
	}
}

func (p *Prompt) SetOnSubmit(callback func(string)) {
	p.onSubmit = callback
}
