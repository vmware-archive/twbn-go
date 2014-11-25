package main

// #therewillbenoise
//
// Rob Szumlakowski
// Allan Baril

import (
	"fmt"
	"time"
	"github.com/cloudfoundry/cli/plugin"
)

type Twbn struct{}

func (t *Twbn) PrintOutput(output []string) {
	for i := len(output) - 1; i >= 0; i -= 1 {
		fmt.Print(output[i])
	}
}

func (t *Twbn) GetNewOutputItems(lastOutput []string, output []string) ([]string) {
	// Find the first item in lastOutput that matches an item in output and return all
	// items in the output before that item
	lastOutputItem := lastOutput[0]
	for i, outputItem := range output {
		if lastOutputItem == outputItem {
			return output[:i]
		}
	}
	return output
}

func (t *Twbn) AnalyzeOutput(lastOutput []string, output []string) {
	if lastOutput == nil {
		return
	}

	newItems := t.GetNewOutputItems(lastOutput, output)
	t.PrintOutput(newItems)
}

func (t *Twbn) Run(cliConnection plugin.CliConnection, args []string) {
	var lastOutput []string
	var output []string
	var err error

	if args[0] == "twbn" {

		if len(args) <= 1 {
			fmt.Println("Usage:\n\tcf twbn APP_NAME")
			return
		}

		appName := args[1]
		fmt.Println("Watching status for application ", appName)

		for {
			output, err = cliConnection.CliCommandWithoutTerminalOutput("events", appName)
			if err != nil {
				fmt.Println("Error: ", err)
				return;
			}

			output = output[2:]
			t.AnalyzeOutput(lastOutput, output)
			time.Sleep(1 * time.Second)
			lastOutput = output
		}
	}
}

func (t *Twbn) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "#therewillbenoise",
		Commands: []plugin.Command{
			{
				Name:     "twbn",
				HelpText: "Get ready to hear some noise.",
			},
		},
	}
}

func main() {
	plugin.Start(new(Twbn))
}

