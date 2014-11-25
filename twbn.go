package main

// #therewillbenoise
//
// Rob Szumlakowski
// Allan Baril

import (
	"fmt"
	"time"
	"strings"
	"os/exec"
	"github.com/cloudfoundry/cli/plugin"
)

type Twbn struct {
	appName string
}

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

func (t *Twbn) Light() {
	cmd := exec.Command("/bin/bash", "light.sh")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running light.sh: ", err)
	}
}

func (t *Twbn) Speech(s string) {
	cmd := exec.Command("/bin/bash", "speech.sh", s)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running speech.sh: ", err)
	}
}

func (t *Twbn) AppStarted() {
	t.Light()
	t.Speech("Application " + t.appName + " started")
}

func (t *Twbn) AppStopped() {
	t.Light()
	t.Speech("Application " + t.appName + " stopped")
}

func (t *Twbn) AppCrashed() {
	t.Light()
	t.Speech("ALERT: Application " + t.appName + " crashed")
}

func (t *Twbn) Instances(s string) {
	i := strings.LastIndex(s, ":") + 1
	numInstances := s[i:]
	t.Light()
	t.Speech("Application " + t.appName + " is now running with " + numInstances + " instances")
}

func (t *Twbn) AnalyzeOutputLine(s string) (bool) {
	if (strings.Contains(s, "state: STARTED")) {
		t.AppStarted()
		return true
	} else if (strings.Contains(s, "state: STOPPED")) {
		t.AppStopped()
		return true
	} else if (strings.Contains(s, "reason: CRASHED")) {
		t.AppCrashed()
		return true
	} else if (strings.Contains(s, "instances: ")) {
		t.Instances(s)
		return true
	}
	return false
}

func (t *Twbn) AnalyzeOutput(lastOutput []string, output []string) {
	if lastOutput == nil {
		return
	}

	newItems := t.GetNewOutputItems(lastOutput, output)
	t.PrintOutput(newItems)

	for i := len(newItems) - 1; i >= 0; i -= 1 {
		isDone := t.AnalyzeOutputLine(newItems[i])
		if isDone {
			break;
		}
	}
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

		t.appName = args[1]
		fmt.Println("Watching status for application ", t.appName)

		for {
			output, err = cliConnection.CliCommandWithoutTerminalOutput("events", t.appName)
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

