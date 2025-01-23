package main

import (
	"hunt/ui"
	"log"
	"os"

	"github.com/joho/godotenv"

	"fyne.io/fyne/v2/app"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	myApp := app.New()
	model := getModel()
	window := ui.NewMainWindow(myApp, model)
	window.Show()
	myApp.Run()
}

// break this off later
func getModel() llms.Model {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	model := os.Getenv("model")
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal("Error creating model", err)
	}
	return llm
}
