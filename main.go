package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"oli/commands"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file")
	}
}

type Command interface {
	Execute(context.Context, []string) error
	Matches([]string) bool
}

var availableCommands = []Command{
	&commands.EchoCommand{},
	&commands.HelpCommand{},
	&commands.ModelsCommand{},
	&commands.QuitCommand{},
}

func findCommands(inputs []string) []Command {
	array := []Command{}

	for _, command := range availableCommands {
		if command.Matches(inputs) {
			array = append(array, command)
		}
	}

	return array
}

func main() {
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		inputs := strings.Split(input, " ")

		var command Command

		matchingCommands := findCommands(inputs)
		matchingCommandsLen := len(matchingCommands)

		if matchingCommandsLen == 0 {
			command = &commands.HelpCommand{
				Model: "qwen2.5-coder:7b",
			}
		} else if matchingCommandsLen == 1 {
			command = matchingCommands[0]
		} else {
			// SKIP
		}

		err := command.Execute(ctx, inputs)
		if err != nil {
			log.Printf("there was a command error: %v", err)
		}
	}
}
