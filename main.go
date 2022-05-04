package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Hardcoded repl commands
	metaCommands := map[string]interface{}{
		".help":  displayHelp,
		".clear": clearScreen,
	}

	statements := map[string]interface{}{
		"select": selectState,
		"insert": insertState,
	}
	// Begin the repl loop
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		if strings.EqualFold(".", string(reader.Bytes()[0])) {
			if command, exists := metaCommands[text]; exists {
				// Call a hardcoded function
				command.(func())()
			} else if strings.EqualFold(".exit", text) {
				// Close the program on the exit command
				return
			} else {
				// Pass the command to the parser
				handleCmd(text)
			}
		} else {
			if statement, exists := statements[text]; exists {
				// Call a hardcoded function
				statement.(func())()
			} else {
				// Pass the command to the parser
				handleCmd(text)
			}
		}

		printPrompt()
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}

// cliName is the name used in the repl prompts
var cliName string = "sqlite"

// printPrompt displays the repl prompt at the start of each loop
func printPrompt() {
	fmt.Print(cliName, "> ")
}

// printUnkown informs the user about invalid commands
func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

// displayHelp informs the user about our hardcoded functions
func displayHelp() {
	fmt.Printf("Welcome to %v! These are the available commands: \n", cliName)
	fmt.Println(".help    - Show available commands")
	fmt.Println(".clear   - Clear the terminal screen")
	fmt.Println(".exit    - Closes your connection to", "sqlite")
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// handleInvalidCmd attempts to recover from a bad command
func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

// handleCmd parses the given commands
func handleCmd(text string) {
	handleInvalidCmd(text)
}

// cleanInput preprocesses input to the db repl
func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func insertState() {
	fmt.Println("INSERT - insert to database")
}

func selectState() {
	fmt.Println("SELECT - select from database")
}
