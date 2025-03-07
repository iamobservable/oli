package commands

import (
	"context"
	"fmt"
	"slices"
	"strings"
)

var VALID_ECHO_ALIASES = []string{"/echo", "/e"}

type EchoCommand struct{}

func (c *EchoCommand) Execute(ctx context.Context, args []string) error {
	message := strings.Join(args[1:], " ")
	// TODO: DEFINE WITH UNICODE AND TRIMRIGHTFUNC

	fmt.Printf("%v\n", strings.Trim(strings.Trim(message, "\n"), "\r"))

	return nil
}

func (c *EchoCommand) Matches(inputs []string) bool {
	return len(inputs) > 0 && slices.Contains(VALID_ECHO_ALIASES, inputs[0])
}
