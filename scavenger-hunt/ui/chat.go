package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Chat struct {
	Container fyne.CanvasObject
	textArea  *widget.RichText
	scroll    *container.Scroll
}

func NewChat() *Chat {
	c := &Chat{
		textArea: widget.NewRichText(),
	}
	c.textArea.Wrapping = fyne.TextWrapWord

	c.scroll = container.NewScroll(c.textArea)
	c.Container = container.NewStack(c.scroll)

	return c
}

func (c *Chat) AddMessage(role, content string) {
	prefix := widget.NewRichTextWithText(role + ": ")
	markdown := widget.NewRichTextFromMarkdown(content)
	newline := widget.NewRichTextWithText("\n\n")

	// Combine all segments
	var allSegments []widget.RichTextSegment
	allSegments = append(allSegments, prefix.Segments...)
	allSegments = append(allSegments, markdown.Segments...)
	allSegments = append(allSegments, newline.Segments...)

	// Add to existing segments
	c.textArea.Segments = append(c.textArea.Segments, allSegments...)
	c.textArea.Refresh()

	c.scroll.ScrollToBottom()
}

func (c *Chat) UpdateThinking(content string) {
	markdown := widget.NewRichTextFromMarkdown(content)
	c.textArea.Segments = append(c.textArea.Segments, markdown.Segments...)
	c.textArea.Refresh()
}
