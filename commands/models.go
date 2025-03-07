package commands

import (
	"context"
	"fmt"
	"oli/services"
	"os"
	"slices"
)

var VALID_MODELS_ALIASES = []string{"/models", "/m"}

type ModelsCommand struct{}

func (c *ModelsCommand) Execute(ctx context.Context, args []string) error {
	ollamaService := services.OllamaService{
		BASE_URL: os.Getenv("OLLAMA_BASE_URL"),
	}

	tagsRequest := services.TagsRequest{}

	tagsResponse, err := ollamaService.Models(&tagsRequest)
	if err != nil {
		return err
	}

	fmt.Println("Models available:")
	for _, model := range tagsResponse.Models {
		fmt.Printf("  %v(%v):%v\n", model.Name, model.Details.Family, model.Details.ParameterSize)
	}

	return nil
}

func (c *ModelsCommand) Matches(inputs []string) bool {
	return len(inputs) > 0 && slices.Contains(VALID_MODELS_ALIASES, inputs[0])
}
