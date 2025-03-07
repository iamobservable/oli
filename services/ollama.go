package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type OllamaService struct {
	BASE_URL string
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

type ChatResponse struct {
	Model      string      `json:"model"`
	CreatedAt  string      `json:"created_at"`
	Message    ChatMessage `json:"message"`
	Done       bool        `json:"done"`
	DoneReason string      `json:"done_response"`
}

type ChatConversation struct {
	Id       string        `json:"id"`
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Path     string        `json:"path"`
}

func GetUniqueChatConversation() (*ChatConversation, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error accessing the user home directory: %v", err)
	}

	path = filepath.Join(path, OLIVER_PATH)

	err = os.MkdirAll(path, os.ModeDir|0755)
	if err != nil {
		return nil, fmt.Errorf("error creating directory: %v", err)
	}

	id := uuid.NewSHA1(uuid.NameSpaceURL, []byte(path)).String()

	cc, err := FindJsonMemoryRecord(path, id)
	if err != nil {
		return &ChatConversation{
			Id:    id,
			Model: "qwen2.5-coder:7b",
			Messages: []ChatMessage{
				NewSystemMessage(`
          You are a helpful, friendly, and direct llm model assistant named Oliver. Here are some details about you:
            - You do not waste time with extra words
            - You do not know who creeated you
            - You typically respond in two sentences or less
            - If you can't respond succinctly, ask for more context before answering
            - You start the first part of your conversations with a greeting

         You are able to help with the following two commands. When a question is not provided, summarize how you can help.
           - /echo
             Echo means you will repeat what is asked
             Example: "echo I repeat things"
           - /help
             Answer the help provided
             Example: "help with golang or python"`),
			},
			Path: path,
		}, nil
	}

	return cc, nil
}

func SaveChatConversation(cc *ChatConversation) error {
	err := SaveJsonMemoryRecord(cc.Path, cc)
	if err != nil {
		return fmt.Errorf("error saving chat conversation: %v", err)
	}

	return nil
}

func NewAssistantMessage(content string) ChatMessage {
	return ChatMessage{
		Role:    "assistant",
		Content: content,
	}
}

func NewSystemMessage(content string) ChatMessage {
	return ChatMessage{
		Role:    "system",
		Content: content,
	}
}

func NewUserMessage(content string) ChatMessage {
	return ChatMessage{
		Role:    "user",
		Content: content,
	}
}

var ollama_base_url = os.Getenv("OLLAMA_BASE_URL")

func (s *OllamaService) Chat(req *ChatRequest) (*bufio.Reader, error) {
	if s.BASE_URL == "" {
		fmt.Print("environment variable missing: ollama_base_url")
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(jsonBody)

	res, err := http.Post(fmt.Sprintf("%v/api/chat", s.BASE_URL), "application/json", requestBody)
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(res.Body), nil
}

// type ListModelsRequest struct {
// 	Model    string        `json:"model"`
// 	Messages []ChatMessage `json:"messages"`
// 	Stream   bool          `json:"stream"`
// }
//
// type ChatResponse struct {
// 	Model      string      `json:"model"`
// 	CreatedAt  string      `json:"created_at"`
// 	Message    ChatMessage `json:"message"`
// 	Done       bool        `json:"done"`
// 	DoneReason string      `json:"done_response"`
// }
//
//
// func *s *OllamaService) Modeels(req *ModelsRequest) (*ModelsResponse, error) {}
