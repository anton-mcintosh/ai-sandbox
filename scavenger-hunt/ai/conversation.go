package ai

import (
	"context"
	"fmt"
	"hunt/types"
	"regexp"

	"github.com/tmc/langchaingo/llms"
)

type Conversation struct {
	model        llms.Model
	conversation string
}

func NewConversation(model llms.Model) *Conversation {
	return &Conversation{
		model:        model,
		conversation: "",
	}
}

func (c *Conversation) HandlePrompt(text string, thinker types.ThinkUpdater) (string, error) {

	c.conversation += fmt.Sprintf("User: %s\n", text)
	response := ""

	ctx := context.Background()
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are the start of a new useful bot."),
		llms.TextParts(llms.ChatMessageTypeHuman, c.conversation),
		llms.TextParts(llms.ChatMessageTypeHuman, text),
	}
	completion, err := c.model.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		response += string(chunk)
		thinker.UpdateThinking(string(chunk))
		return nil
	}))
	if err != nil {
		fmt.Println("Error: ", err)
	}

	re := regexp.MustCompile(`(?s)<think>.*?</think>`)
	finalResponse := re.ReplaceAllString(response, "")
	c.conversation += fmt.Sprintf("Assistant: %s\n", finalResponse)
	_ = completion
	return finalResponse, nil
}

// UpdateThinking implements the ThinkUpdater interface
func (c *Conversation) UpdateThinking(status string) {
	// Add any thinking status update logic here
}
