package commands

import (
	"context"
	"fmt"
	"os"
	"slices"
)

var VALID_MODELS_ALIASES = []string{"/models", "/m"}

type ModelsCommand struct{}

func (c *ModelsCommand) Execute(ctx context.Context, args []string) error {
	fmt.Print("Until next time!\n")
	os.Exit(0)

	return nil
}

func (c *ModelsCommand) Matches(inputs []string) bool {
	return len(inputs) > 0 && slices.Contains(VALID_MODELS_ALIASES, inputs[0])
}
