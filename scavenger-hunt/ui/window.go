package ui

import (
	"fmt"
	"hunt/ai"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/tmc/langchaingo/llms"
)

type MainWindow struct {
	Window       fyne.Window
	Functions    *Functions 
	Chat         *Chat     
	Prompt       *Prompt    
	Think        *Think     
	Conversation *ai.Conversation
}

func NewMainWindow(app fyne.App, model llms.Model) *MainWindow {
	w := &MainWindow{
		Window:       app.NewWindow("LLM Scavenger Hunt"),
		Conversation: ai.NewConversation(model),
	}

	w.Functions = NewFunctions()       
	w.Chat = NewChat()                 
	w.Prompt = NewPrompt()             
	w.Think = NewThink(w.Conversation) 

	// retrieves text from the prompt window and sends it to handlePrompt
	w.Prompt.SetOnSubmit(func(text string) {
		w.handlePrompt(text)
	})

	w.setupUI()
	return w
}

func (w *MainWindow) setupUI() {
	middle := container.NewVSplit(
		w.Chat.Container,
		w.Prompt.Container,
	)
	middle.SetOffset(0.8) 

	leftSplit := container.NewHSplit(
		w.Functions.Container, 
		middle,               
	)
	leftSplit.SetOffset(0.2) 

	rightSplit := container.NewHSplit(
		leftSplit,         
		w.Think.Container, 
	)
	rightSplit.SetOffset(0.8) /

	w.Window.SetContent(rightSplit)
	w.Window.Resize(fyne.NewSize(1200, 800)) 
}

func (w *MainWindow) handlePrompt(text string) {

	w.Chat.AddMessage("User", text)
	response, err := w.Conversation.HandlePrompt(text, w.Think)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	w.Chat.AddMessage("Assistant", response)
}

func (w *MainWindow) Show() {
	w.Window.Show()
}
