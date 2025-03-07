package commands

import (
	"context"
	"fmt"
	"os"
	"slices"
)

var VALID_QUIT_ALIASES = []string{"/quit", "/exit", "/q"}

type QuitCommand struct{}

func (c *QuitCommand) Execute(ctx context.Context, args []string) error {
	fmt.Print("Until next time!\n")
	os.Exit(0)

	return nil
}

func (c *QuitCommand) Matches(inputs []string) bool {
	return len(inputs) > 0 && slices.Contains(VALID_QUIT_ALIASES, inputs[0])
}
