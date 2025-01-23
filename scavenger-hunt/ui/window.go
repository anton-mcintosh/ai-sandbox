package ui

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/tmc/langchaingo/llms"
)

type MainWindow struct {
	Window       fyne.Window
	Functions    *Functions // Left sidebar
	Chat         *Chat      // Middle top
	Prompt       *Prompt    // Middle bottom
	Think        *Think     // Right sidebar
	Model        llms.Model
	Conversation string
}

func NewMainWindow(app fyne.App, model llms.Model) *MainWindow {
	w := &MainWindow{
		Window: app.NewWindow("LLM Scavenger Hunt"),
		Model:  model,
	}

	w.Functions = NewFunctions() // Available LLM functions
	w.Chat = NewChat()           // Chat history display
	w.Prompt = NewPrompt()       // User input and send button
	w.Think = NewThink()         // LLM thinking process

	// retrieves text from the prompt window and sends it to handlePrompt
	w.Prompt.SetOnSubmit(func(text string) {
		w.handlePrompt(text)
	})

	w.setupUI()
	return w
}

func (w *MainWindow) setupUI() {
	// Create middle section (chat + prompt)
	middle := container.NewVSplit(
		w.Chat.Container,
		w.Prompt.Container,
	)
	middle.SetOffset(0.8) // Chat takes 80% height, prompt 20%

	// Create main layout with sidebars
	leftSplit := container.NewHSplit(
		w.Functions.Container, // Left sidebar
		middle,                // Middle section
	)
	leftSplit.SetOffset(0.2) // Functions takes 20% width

	rightSplit := container.NewHSplit(
		leftSplit,         // Combined left side + middle
		w.Think.Container, // Right sidebar
	)
	rightSplit.SetOffset(0.8) // Main content takes 80% width, Think takes 20%

	w.Window.SetContent(rightSplit)
	w.Window.Resize(fyne.NewSize(1200, 800)) // Set initial window size
}

func (w *MainWindow) handlePrompt(text string) {
	// Show user message
	w.Chat.AddMessage("User", text)
	w.Conversation += fmt.Sprintf("User: %s\n", text)
	response := ""

	// Get LLM response
	ctx := context.Background()
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are the start of a new useful bot."),
		llms.TextParts(llms.ChatMessageTypeHuman, w.Conversation),
		llms.TextParts(llms.ChatMessageTypeHuman, text),
	}
	completion, err := w.Model.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		response += string(chunk)
		return nil
	}))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Show LLM response
	w.Chat.AddMessage("Assistant", response)
	w.Conversation += fmt.Sprintf("Assistant: %s\n", response)
	_ = completion
}

func (w *MainWindow) Show() {
	w.Window.Show()
}
