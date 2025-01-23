package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Chat struct {
	Container fyne.CanvasObject
	textArea  *widget.Entry
	// Add other fields needed
}

func NewChat() *Chat {
	c := &Chat{
		textArea: widget.NewMultiLineEntry(),
	}
	c.textArea.Wrapping = fyne.TextWrapWord
	c.textArea.MultiLine = true

	// Set minimum size for the text area
	c.textArea.Resize(fyne.NewSize(800, 600))
	c.textArea.MinSize()

	// Create scrollable container
	scroll := container.NewScroll(c.textArea)
	c.Container = container.NewStack(scroll)

	return c
}

// AddMessage adds a new message synchronously
func (c *Chat) AddMessage(role, content string) {
	message := fmt.Sprintf("%s: %s\n\n", role, content)
	currentText := c.textArea.Text
	c.textArea.SetText(currentText + message)

	// Scroll to bottom after new content
	c.textArea.CursorRow = len(c.textArea.Text)
	c.textArea.Refresh()
}

// UpdateThinking updates the thinking process in real-time
func (c *Chat) UpdateThinking(content string) {
	c.textArea.SetText(c.textArea.Text + content)
	c.textArea.CursorRow = len(c.textArea.Text)
	c.textArea.Refresh()
}
