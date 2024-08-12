package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var cliName string = "Pokedex"

func printPrompt() {
	fmt.Print(cliName, "> ")
}

func invalidCommandError(text string) {
	fmt.Print(text, "Command not found")
}

func displayHelp() {
	fmt.Printf(
		"Welcome to %v! These are the available commands: \n",
		cliName,
	)
	fmt.Println(".help    - Show available commands")
	fmt.Println(".clear   - Clear the terminal screen")
	fmt.Println(".exit    - Closes your connection to ", cliName)
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func handleInvalidCmd(text string) {
	defer invalidCommandError(text)
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

// handleCmd parses the given commands
func handleCmd(text string) {
	handleInvalidCmd(text)
}

func main() {
	// Hardcoded repl commands
	commands := map[string]interface{}{
		".help":  displayHelp,
		".clear": clearScreen,
	}
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if command, exists := commands[text]; exists {
			// Call a hardcoded function
			command.(func())()
		} else if strings.EqualFold(".exit", text) {
			// Close the program on the exit command
			return
		} else {
			// Pass the command to the parser
			handleCmd(text)
		}
		printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}
