package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"oli/services"
	"os"
	"slices"
	"strings"
)

var VALID_HELP_ALIASES = []string{"/help", "/h"}

type HelpCommand struct{}

func (c *HelpCommand) Execute(ctx context.Context, args []string) error {
	cc, err := services.GetUniqueChatConversation()
	if err != nil {
		return fmt.Errorf("error getting chat converstation: %v", err)
	}

	if len(args) == 0 {
		cc.Messages = append(cc.Messages, services.NewUserMessage("How can you help me?"))
	} else {
		cc.Messages = append(cc.Messages, services.NewUserMessage(fmt.Sprintf("Provide help on %v", strings.Join(args, " "))))
	}

	chatRequest := services.ChatRequest{
		Model:    cc.Model,
		Messages: cc.Messages,
		Stream:   true,
	}

	ollamaService := services.OllamaService{
		BASE_URL: os.Getenv("OLLAMA_BASE_URL"),
	}

	reader, err := ollamaService.Chat(&chatRequest)
	if err != nil {
		return err
	}

	content := []string{}

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatalf("reading bytes failed: %v", err)
		}

		var cr services.ChatResponse

		err = json.Unmarshal([]byte(line), &cr)
		if err != nil {
			log.Fatalf("failed to unmarshal json: %v", err)
		}

		fmt.Print(cr.Message.Content)
		content = append(content, cr.Message.Content)

		if cr.Done {
			assistentMessage := services.NewAssistantMessage(strings.Join(content, ""))
			cc.Messages = append(cc.Messages, assistentMessage)

			err := services.SaveChatConversation(cc)
			if err != nil {
				return fmt.Errorf("error saving chat conversation: %v", err)
			}

			fmt.Printf("%v", "\n\n")
			break
		}
	}

	return nil
}

func (c *HelpCommand) Matches(inputs []string) bool {
	return len(inputs) > 0 && slices.Contains(VALID_HELP_ALIASES, inputs[0])
}
